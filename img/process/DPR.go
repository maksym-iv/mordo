package process

// DPR - simply enlarge image according to ratio
// Uses BimgResize()
func (i *Image) DPR(ratio float64) error {
	if err := i.img.Resize(ratio); err != nil {
		return err
	}

	return nil
}
