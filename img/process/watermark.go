package process

import (
	"math"

	"github.com/davidbyttow/govips/v2/vips"
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
	imgOverlay, err := vips.NewImageFromFile(config.WatermarkConfig.Path)
	if err != nil {
		return err
	}

	// scale watermark according to config.WatermarkConfig.Scale.
	// TODO: put `imgOverlay.Resize` out of `if` scope
	if scale != 0 {
		s := math.Round(float64(config.WatermarkConfig.Scale * scale))
		imgOverlay.Resize(s, vips.KernelAuto)
	} else {
		s := math.Round(float64(config.WatermarkConfig.Scale))
		imgOverlay.Resize(s, vips.KernelAuto)
	}

	// TODO: Put this struct somewhere, name it `exportParamsHQ` and re-use it.
	exportParams := &vips.ExportParams{
		Quality:     100,
		Compression: 0,
		Lossless:    true,
		Format:      imgOverlay.Format(),
	}

	imgOverlayBuff, _, err := imgOverlay.Export(exportParams)
	if err != nil {
		return err
	}
	imgOverlay, err = vips.NewImageFromBuffer(imgOverlayBuff)
	if err != nil {
		return err
	}

	// Define X, Y of for watermark position
	imgOverlayX := setX(i.img.Width(), imgOverlay.Width(), xString)
	imgOverlayY := setY(i.img.Height(), imgOverlay.Height(), yString)

	if err := i.img.Composite(imgOverlay, vips.BlendModeOver, imgOverlayX, imgOverlayY); err != nil {
		return err
	}

	return nil
}
