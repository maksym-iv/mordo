package process

import (
	"github.com/davidbyttow/govips/v2/vips"

	c "github.com/xmackex/mordo/config"
)

var (
	config *c.Config
)

func init() {
	cfg, err := c.NewConfig()
	if err != nil {
		panic(err)
	}
	config = cfg
}

type Image struct {
	img *vips.ImageRef
}

type applied struct {
	width  int
	height int
}

// New - just create new *bimg.Image
func New(buff []byte) (*Image, error) {
	img, err := vips.NewImageFromBuffer(buff)
	if err != nil {
		return nil, err
	}
	i := &Image{
		img: img,
	}
	return i, nil
}

// Process - process *bimg.Image with *bimg.Options
func (i *Image) Process() ([]byte, *string, error) {
	exportParams := &vips.ExportParams{
		Quality:     config.ImageConfig.Quality,
		Compression: config.ImageConfig.Compression,
		Lossless:    config.ImageConfig.Lossless,
		Format:      i.img.Format(),
	}

	buff, imgExt, err := i.img.Export(exportParams)
	if err != nil {
		return nil, nil, err
	}

	ext := imgExt.Format.FileExt()

	return buff, &ext, nil
}
