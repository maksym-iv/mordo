package main

import (
	"encoding/base64"
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"

	c "github.com/xmackex/mordo/config"
	p "github.com/xmackex/mordo/img/process"
	"github.com/xmackex/mordo/io/s3"
)

var (
	log    *zap.SugaredLogger
	config *c.Config
)

func init() {
	var err error
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	log = logger.Sugar()

	config, err = c.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func response(b string, s int, h map[string]string) *events.APIGatewayProxyResponse {
	h["Cache-Control"] = "no-cache"
	var r events.APIGatewayProxyResponse
	r.Body = b
	r.StatusCode = s
	r.Headers = h
	return &r
}

func timer(timeStart time.Time, where string) {
	timeNow := time.Now()
	timePassed := timeNow.Sub(timeStart)
	log.Debugf("Passed after %s: %v", where, timePassed)
}

func process(image string, qs map[string]string, stageVars map[string]string) (*events.APIGatewayProxyResponse, *gwError) {
	timeStart := time.Now()

	vips.Startup(&vips.Config{
		ConcurrencyLevel: 4,
		MaxCacheFiles:    100,
		MaxCacheMem:      config.MaxMem,
		MaxCacheSize:     500,
		ReportLeaks:      false,
		CacheTrace:       false,
		CollectStats:     false,
	})

	qss, err := newQs(image, qs)
	newPath := qss.hashPath()

	if err != nil {
		log.Error(err)
		e := newErr(err.Error())
		r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
		return r, e
	}
	if qss.Image == "" {
		e := newErr("Image key for processing was not set")
		log.Error(err)
		r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
		return r, e
	}

	s3Img, err := s3.NewImg(stageVars["s3_bucket"], qss.Image, stageVars["s3_bucket_region"])
	if err != nil {
		log.Error(err)
		e := newErr(err.Error())
		r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
		return r, e
	}

	// Check if we need process image
	if ok, err := s3Img.IsExist(config.AWSConfig.S3BucketProcessed, newPath); err != nil {
		log.Error(err)
		e := newErr(err.Error())
		r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
		return r, e
	} else if ok && !config.Debug {
		s3Img.UpdatePath(newPath)
		s3Img.UpdateBucket(config.AWSConfig.S3BucketProcessed)
		s3Img.Read()
		s3Img.SetContentType()
		r := &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type":  s3Img.ContentType,
				"Cache-Control": fmt.Sprintf("max-age=%d", config.Cache.MaxAge),
			},
			Body:            base64.StdEncoding.EncodeToString(s3Img.Buff),
			IsBase64Encoded: true,
		}
		return r, nil
	}

	// Check if original image exist
	if ok, err := s3Img.IsExist(s3Img.Bucket, s3Img.Path); err != nil {
		log.Error(err)
		e := newErr(err.Error())
		r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
		return r, e
	} else if !ok {
		e := newErr("Initial image doesn't exist")
		log.Error(err)
		// TODO: maybe 404?
		r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
		return r, e
	}

	timeStart = time.Now()

	// START transform
	s3Img.Read()

	timer(timeStart, "Read") // DEBUGG
	img, err := p.New(s3Img.Buff)
	if err != nil {
		log.Error(err)
		e := newErr(err.Error())
		r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
		return r, e
	}

	if qss.Crop.Width != 0 && qss.Crop.Height != 0 {
		if err := img.Crop(qss.Crop.Left, qss.Crop.Top, qss.Crop.Width, qss.Crop.Height); err != nil {
			log.Error(err)
			e := newErr(err.Error())
			r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
			return r, e
		}
	}

	timer(timeStart, "Crop") // DEBUGG
	if qss.Resize.Width != 0 {
		if err := img.Resize("width", qss.Resize.Width); err != nil {
			log.Error(err)
			e := newErr(err.Error())
			r := response(e.ErrorJSON(), 404, map[string]string{"Content-Type": "application/json"})
			return r, e
		}
	} else if qss.Resize.Height != 0 {
		if err := img.Resize("height", qss.Resize.Height); err != nil {
			log.Error(err)
			e := newErr(err.Error())
			r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
			return r, e
		}
	}

	timer(timeStart, "Resize") // DEBUGG
	if qss.DPR != 0 {
		img.DPR(qss.DPR)
	}

	timer(timeStart, "DPR") // DEBUGG
	if qss.Sharpen {
		log.Infof("Will sharpen: %s", s3Img.Path)
		// http://jcupitt.github.io/libvips/API/current/libvips-convolution.html#vips-sharpen
		if err := img.Sharpen(1, 2, 3); err != nil {
			log.Error(err)
			e := newErr(err.Error())
			r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
			return r, e
		}
	}

	timer(timeStart, "Sharpen") // DEBUGG
	if qss.Watermark.WX != "" && qss.Watermark.WY != "" && config.WatermarkEnabled {
		var err error
		if err = img.Watermark(qss.Watermark.WX, qss.Watermark.WY, qss.Watermark.WS); err != nil {
			log.Error(err)
			e := newErr(err.Error())
			r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
			return r, e
		}
	}

	// END transform
	timer(timeStart, "Watermark") // DEBUGG
	buf, _, err := img.Process()
	if err != nil {
		log.Error(err)
		e := newErr(err.Error())
		r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
		return r, e
	}
	timer(timeStart, "Process") // DEBUGG

	s3Img.UpdatePath(newPath)
	s3Img.UpdateBuff(buf)
	timer(timeStart, "UpdateBuff") // DEBUGG
	if err := s3Img.Write(); err != nil {
		log.Error(err)
		e := newErr(err.Error())
		r := response(e.ErrorJSON(), 500, map[string]string{"Content-Type": "application/json"})
		return r, e
	}
	timer(timeStart, "Write") // DEBUGG

	r := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":  s3Img.ContentType,
			"Cache-Control": fmt.Sprintf("max-age=%d", config.Cache.MaxAge),
		},
		Body:            base64.StdEncoding.EncodeToString(s3Img.Buff),
		IsBase64Encoded: true,
	}
	return r, nil

}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// https://mg66xt11sl.execute-api.us-west-1.amazonaws.com/dev-dev-mordo-images/process?image=tests/read/origin_2k.jpg

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Infof("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no name is provided in the HTTP request body, throw an error
	image, ok := request.PathParameters["image"]
	if !ok {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("No image was defined")
	}

	response, err := process(image, request.QueryStringParameters, request.StageVariables)
	if err != nil {
		log.Error(err)
	}

	return *response, nil

}

func main() {
	lambda.Start(handler)
}
