package process

import (
	"fmt"

	"github.com/davidbyttow/govips/v2/vips"
)

// Resize - resize image to height or width according to size with govips lib
func (i *Image) Resize(dimension string, size int) error {
	var scale float64 = 1

	switch dimension {
	case "width", "w":
		scale = float64(size) / float64(i.img.Width())
	case "height", "h":
		scale = float64(size) / float64(i.img.Height())
	default:
		err := fmt.Errorf("Invalid dimension. Can be \"h/height\" or \"w/width\"")
		return err
	}

	if err := i.img.Resize(scale, vips.KernelAuto); err != nil {
		return err
	}

	return nil
}
