package process

import ()

// Crop - crop
func (i *Image) Crop(left, top, width, height int) error {
	// width, height can't be more than image Width, Height
	if width > i.img.Width() {
		width = i.img.Width()
	}
	if height > i.img.Height() {
		height = i.img.Height()
	}

	// set left and top to crop img.With - crop.Width, img.Length - crop.Length if left, tom more than image width or length
	if left+width > i.img.Width() {
		left = i.img.Width() - width
	}
	if top+height > i.img.Height() {
		top = i.img.Height() - height
	}

	if err := i.img.ExtractArea(left, top, width, height); err != nil {
		return err
	}

	return nil
}
