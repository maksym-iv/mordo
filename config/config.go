package config

import (
	"fmt"
	"strings"
	// "strconv"
	"github.com/kelseyhightower/envconfig"
	// "github.com/xmackex/resizer/helpers"
)

const (
	platforms = "aws"
)

type Config struct {
	Debug            bool      `default:"False"`
	Platform         string    `required:"true"`
	MaxMem           int       `default:"128"`
	Cache            Cache     `ignored:"true"`
	AWSConfig        AWS       `ignored:"true"`
	WatermarkEnabled bool      `default:"False" envconfig:"watermark"`
	WatermarkConfig  Watermark `ignored:"true"`
	ImageConfig      Image     `ignored:"true"`
	// CloudFrontID string `required:"false"`
}

// Cache - definec cacheing headers
type Cache struct {
	MaxAge int `required:"true"`
}

// AWS config struct
type AWS struct {
	S3BucketProcessed string `required:"true" split_words:"true"`
}

// Watermark config struct
type Watermark struct {
	Path  string  `required:"true" split_words:"true"`
	Scale float64 `split_words:"true" default:"0.15"`
}

// Image config struct
type Image struct {
	Quality     int  `required:"false" default:"90"`
	Lossless    bool `required:"false" default:"false"`
	Compression int  `required:"false" default:"6"`
	Enlarge     bool `required:"false" default:"false"`
}

func platformSupported(p string) bool {
	pa := strings.Split(platforms, ",")
	for i := range pa {
		if pa[i] == p {
			return true
		}
	}
	return false
}

// NewConfig - create new config from env
func NewConfig() (*Config, error) {
	config := &Config{}
	cache := &Cache{}
	awsConfig := &AWS{}
	watermarkConfig := &Watermark{}
	imageConfig := &Image{}

	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	config.Platform = strings.ToLower(config.Platform)

	if ps := platformSupported(config.Platform); !ps {
		err := fmt.Errorf("Now valid platform defined. Supported platforms: %s", strings.Join(strings.Split(platforms, ","), ", "))
		return nil, err
	}

	// Select and parse platform
	switch config.Platform {
	case "aws":
		err := envconfig.Process("aws", awsConfig)
		if err != nil {
			return nil, err
		}
	default:
		err := fmt.Errorf("Error while parsing config. Call the developer")
		return nil, err
	}

	if err := envconfig.Process("cache", cache); err != nil {
		err := fmt.Errorf("Cache setting was not set")
		return nil, err
	}

	// Watermark settings
	if config.WatermarkEnabled {
		if err := envconfig.Process("watermark", watermarkConfig); err != nil {
			// err := fmt.Errorf("Watermark enabled and WATERMARK_PATH was not set")
			return nil, err
		}
	}

	// Image settings
	if err := envconfig.Process("image", imageConfig); err != nil {
		err := fmt.Errorf("Image settings was not set")
		return nil, err
	}

	config.Cache = *cache
	config.AWSConfig = *awsConfig
	config.WatermarkConfig = *watermarkConfig
	config.ImageConfig = *imageConfig

	return config, nil
}
