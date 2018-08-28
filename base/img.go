package base

// Img represents image object holds image data.
type Img struct {
	Path    string
	DirPath string
	Buff    *[]byte
}

// UpdateBuff - update image with new buffer
func (i *Img) UpdateBuff(b *[]byte) {
	i.Buff = b
}

// UpdatePath - update image with new buffer
func (i *Img) UpdatePath(p string) {
	i.Path = p
}
