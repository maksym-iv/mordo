package config

import (
	"os"
	"reflect"
	"testing"
)

func Test_NewConfig(t *testing.T) {
	testCases := []struct {
		name string
		env  map[string]string
		want *Config
	}{
		{
			name: "AWS",
			env: map[string]string{
				"DEBUG":                   "true",
				"PLATFORM":                "AWS",
				"MAXMEM":                  "128",
				"CACHE_MAXAGE":            "3600",
				"AWS_S3_BUCKET_PROCESSED": "foo-bar",
				"WATERMARK":               "true",
				"WATERMARK_PATH":          "foo/bar/baz/watermark",
				"WATERMARK_SCALE":         "0.15",
				"IMAGE_QUALITY":           "90",
				"IMAGE_COMPRESSION":       "0",
				"IMAGE_LOSSLESS":          "true",
				"IMAGE_ENLARGE":           "TRUE",
			},
			want: &Config{
				Debug:    true,
				Platform: "aws",
				MaxMem:   128,
				Cache:    Cache{MaxAge: 3600},
				AWSConfig: AWS{
					S3BucketProcessed: "foo-bar",
				},
				WatermarkEnabled: true,
				WatermarkConfig: Watermark{
					Path:  "foo/bar/baz/watermark",
					Scale: 0.15,
				},
				ImageConfig: Image{
					Quality:     90,
					Lossless:    true,
					Compression: 0,
					Enlarge:     true,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Parallel()
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.env {
				if err := os.Setenv(k, v); err != nil {
					t.Fatal(err)
				}
			}

			if got, err := NewConfig(); err != nil {
				t.Fatal(err)
			} else if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Config no equal.\nGot: %+v'\nWant: %+v'", got, tc.want)
				t.Errorf("Config no equal.\nGot: %+v'\nWant: %+v'", got.WatermarkConfig, tc.want.WatermarkConfig)
				t.Errorf("Config no equal.\nGot: %+v'\nWant: %+v'", got.AWSConfig, tc.want.AWSConfig)
				t.Errorf("Config no equal.\nGot: %+v'\nWant: %+v'", got.ImageConfig, tc.want.ImageConfig)
			}
		})
	}
}

func Benchmark_NewConfig(b *testing.B) {
	testCases := []struct {
		name string
		env  map[string]string
	}{
		{
			name: "AWS",
			env: map[string]string{
				"DEBUG":    "true",
				"PLATFORM": "AWS",
				// "AWS_S3_BUCKET_NAME":       "foo",
				"AWS_S3_BUCKET_KEY_PREFIX": "/foo/bar/baz",
				"AWS_CLOUDFRONT_DOMAIN":    "foo.example.com",
			},
		},
	}

	b.ReportAllocs()
	b.SetBytes(2)
	b.ResetTimer()

	for _, tc := range testCases {
		b.StopTimer()
		for k, v := range tc.env {
			if err := os.Setenv(k, v); err != nil {
				b.Fatal(err)
			}
		}
		b.StartTimer()
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if _, err := NewConfig(); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
