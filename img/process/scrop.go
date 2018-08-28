package process

import (
	"github.com/davidbyttow/govips/pkg/vips"
)

// SCrop - smart crop
func (i *Image) SCrop(width, height int) error {
	// interesting

	// http://jcupitt.github.io/libvips/API/current/libvips-conversion.html#VIPS-INTERESTING-ATTENTION:CAPS
	optVipsInterestingAttention := vips.InputInt("interesting", 3)

	if err := i.img.Smartcrop(width, height, optVipsInterestingAttention); err != nil {
		return err
	}

	return nil
}
