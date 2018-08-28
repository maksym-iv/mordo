package process

import (
	"math"

	"github.com/davidbyttow/govips/pkg/vips"
)

func setX(main int, sub int, side string) int {
	switch side {
	case "left":
		return 0
	case "right":
		if pos := main - sub; pos >= 0 {
			return pos
		}
		return 0
	case "center":
		if pos := main/2 - sub/2; pos >= 0 {
			return pos
		}
		return 0
	default:
		return main - sub
	}
}

func setY(main int, sub int, side string) int {
	switch side {
	case "top":
		return 0
	case "bottom":
		if pos := main - sub; pos >= 0 {
			return pos
		}
		return 0
	case "center":
		if pos := main/2 - sub/2; pos >= 0 {
			return pos
		}
		return 0
	default:
		return main - sub
	}
}

// Watermark - set and return watermark *bimg.Options
// Watermark position is set according to xString and yString
// Possible watermark positions:
// 	x: left, right, center
// 	y: top, bottom, center
func (i *Image) Watermark(xString string, yString string, scale float64) error {
	// var err error
	// var timeNow time.Time
	// var timePassed time.Duration
	// timeStart := time.Now()
	// timeNow = time.Now()
	// timePassed = timeNow.Sub(timeStart)
	// log.Debugf("Passed after validateQS: %v", timePassed)

	imgSub, err := vips.NewImageFromFile(config.WatermarkConfig.Path)
	if err != nil {
		return err
	}

	imgSubTransform := vips.NewTransform().Quality(100).Compression(0).Lossless()
	imgSubTransform.Image(imgSub)

	// scale watermark according to config.WatermarkConfig.Scale.
	if scale != 0 {
		imgSubWidth := int(math.Round(float64(i.img.Width()) * config.WatermarkConfig.Scale * scale))
		imgSubTransform.ResizeWidth(imgSubWidth)
	} else {
		imgSubWidth := int(math.Round(float64(i.img.Width()) * config.WatermarkConfig.Scale))
		imgSubTransform.ResizeWidth(imgSubWidth)
	}

	imgSubBuff, _, err := imgSubTransform.Apply()
	if err != nil {
		return err
	}
	imgSub, err = vips.NewImageFromBuffer(imgSubBuff)
	if err != nil {
		return err
	}

	// Define X, Y of for watermark position
	imgSubX := setX(i.img.Width(), imgSub.Width(), xString)
	imgSubY := setY(i.img.Height(), imgSub.Height(), yString)
	err = imgSub.Embed(imgSubX, imgSubY, i.img.Width(), i.img.Height())

	i.img.Composite(imgSub, vips.BlendModeOver)

	return nil
}
