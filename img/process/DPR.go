package process

import (
	"github.com/davidbyttow/govips/v2/vips"
)

// DPR - simply enlarge image according to ratio
// Uses imgResize()
func (i *Image) DPR(ratio float64) error {
	if err := i.img.Resize(ratio, vips.KernelAuto); err != nil {
		return err
	}

	return nil
}
