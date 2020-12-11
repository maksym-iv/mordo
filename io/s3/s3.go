package s3

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/xmackex/mordo/base"
	c "github.com/xmackex/mordo/config"
)

var (
	config *c.Config
)

func init() {
	var err error
	config, err = c.NewConfig()
	if err != nil {
		panic(err)
	}
}

type S3Img struct {
	*base.Img
	ContentType string

	Region string
	Bucket string
	sess   *session.Session
	svc    *s3.S3
}

// NewImg - Init image obj with image path, image path basename, image name.
func NewImg(bucket string, imgpath string, region string) (*S3Img, error) {
	img := &S3Img{}
	img.Img = &base.Img{}
	img.Path = path.Clean(imgpath)
	img.Bucket = bucket
	img.Region = region
	img.DirPath = path.Dir(img.Path)

	// TODO: Deal somehow with it. Make it better..
	var awsConfig *aws.Config
	credsFile := fmt.Sprintf("%s/.aws/credentials", os.Getenv("HOME"))
	// In case of shared file creds. For local test purposes
	if _, err := os.Stat(credsFile); !os.IsNotExist(err) {
		creds := credentials.NewCredentials(&credentials.SharedCredentialsProvider{
			Profile:  "mordo",
			Filename: "",
		})
		awsConfig = &aws.Config{
			Credentials: creds,
		}
	} else {
		awsConfig = &aws.Config{}
	}

	img.sess = session.Must(session.NewSession(awsConfig))

	img.svc = s3.New(img.sess, aws.NewConfig().WithRegion(region))

	return img, nil
}

// IsExist - check if imgpath exist in bucket
func (i *S3Img) IsExist(bucket string, imgpath string) (bool, error) {
	headInput := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(imgpath),
	}

	if head, err := i.svc.HeadObject(headInput); head.ContentLength == nil {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf(err.Error())
	}

	return true, nil
}

// SetContentType - set object Content-Type according to i.Buff
func (i *S3Img) SetContentType() {
	ct := http.DetectContentType(i.Buff)
	i.ContentType = ct
}

// UpdateBucket - set object Content-Type according to i.Buff
func (i *S3Img) UpdateBucket(bucket string) {
	i.Bucket = bucket
}

// Read - read image from S3 and update i.Buff with image buffer
func (i *S3Img) Read() error {
	// sess = session.Must(session.NewSession())
	// svc = s3.New(sess, aws.NewConfig().WithRegion(i.Region))

	input := &s3.GetObjectInput{
		Bucket: aws.String(i.Bucket),
		Key:    aws.String(i.Path),
	}

	result, err := i.svc.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				return fmt.Errorf("%v, %v", s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				return fmt.Errorf(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return fmt.Errorf(err.Error())
		}
	}

	emptyBuf := make([]byte, int(*result.ContentLength))
	i.UpdateBuff(emptyBuf)

	_, err = io.ReadFull(result.Body, i.Buff)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

// Write - write file from buffer i.Buff
func (i *S3Img) Write() error {
	tags := "processed=true"
	input := &s3.PutObjectInput{
		Body:        aws.ReadSeekCloser(bytes.NewReader(i.Buff)),
		Tagging:     &tags,
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(config.AWSConfig.S3BucketProcessed),
		Key:         aws.String(i.Path),
		ContentType: aws.String(i.ContentType),
	}

	_, err := i.svc.PutObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return fmt.Errorf(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return fmt.Errorf(err.Error())
		}
	}

	headInput := &s3.HeadObjectInput{
		Bucket: aws.String(config.AWSConfig.S3BucketProcessed),
		Key:    aws.String(i.Path),
	}

	if err := i.svc.WaitUntilObjectExists(headInput); err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
