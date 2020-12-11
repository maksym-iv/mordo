package process

// Sharpen - set and return sharpen *bimg.Options
func (i *Image) Sharpen(radius float64, X1, M2 float64) error {

	// http://jcupitt.github.io/libvips/API/current/libvips-convolution.html#vips-sharpen
	// if err := i.img.Sharpen(sigma, x1, y2, y3, m1, m2); err != nil {
	if err := i.img.Sharpen(radius, X1, M2); err != nil {
		return err
	}

	return nil
}
