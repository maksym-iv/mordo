package process

import (
	"github.com/davidbyttow/govips/pkg/vips"
)

// Sharpen - set and return sharpen *bimg.Options
func (i *Image) Sharpen(radius float64, X1 float64, Y2 float64, Y3 float64, M1 float64, M2 float64) error {

	// http://jcupitt.github.io/libvips/API/current/libvips-convolution.html#vips-sharpen
	// sharpen := &bimg.Sharpen{
	// 	Radius: radius,
	// 	X1:     X1,
	// 	Y2:     Y2,
	// 	Y3:     Y3,
	// 	M1:     M1,
	// 	M2:     M2,
	// }

	// sigma == 0.5
	// x1 == 2
	// y2 == 10         (don't brighten by more than 10 L*)
	// y3 == 20         (can darken by up to 20 L*)
	// m1 == 0          (no sharpening in flat areas)
	// m2 == 3          (some sharpening in jaggy areas)

	// sigma := vips.NewOption("sigma", C.G_TYPE_ULONG, true, func(gv *C.GValue))
	// sigma := vips.NewOption("sigma", C.G_TYPE_FLOAT, true, nil)
	sigma := vips.InputDouble("sigma", radius)
	x1 := vips.InputDouble("x1", X1)
	y2 := vips.InputDouble("y2", Y2)
	y3 := vips.InputDouble("y3", Y3)
	m1 := vips.InputDouble("m1", M1)
	m2 := vips.InputDouble("m2", M2)

	if err := i.img.Sharpen(sigma, x1, y2, y3, m1, m2); err != nil {
		return err
	}
	// i.opt.Sharpen = *sharpen

	return nil
}
