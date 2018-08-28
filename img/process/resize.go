package process

import (
	"fmt"
	"github.com/davidbyttow/govips/pkg/vips"
)

// Resize - resize image to height or width according to size with govips lib
func (i *Image) Resize(dimension string, size int) error {
	tr := vips.NewTransform().Quality(100).Compression(0).Lossless()
	tr.Kernel(vips.KernelLanczos3)
	tr.Format(i.img.Format())
	tr.Image(i.img)

	switch dimension {
	case "width":
		tr.ResizeWidth(size)
	case "w":
		tr.ResizeWidth(size)
	case "height":
		tr.ResizeHeight(size)
	case "h":
		tr.ResizeHeight(size)
	default:
		err := fmt.Errorf("Invalid dimension. Can be \"h/height\" or \"w/width\"")
		return err
	}

	buff, _, err := tr.Apply()
	if err != nil {
		return err
	}

	if i.img, err = vips.NewImageFromBuffer(buff); err != nil {
		return err
	}

	return nil
}
