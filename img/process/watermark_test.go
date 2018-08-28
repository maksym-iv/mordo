// +build !test

package process

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	// "github.com/xmackex/bimg"
	// // "gopkg.in/h2non/bimg.v1"

	"gitlab.com/doctor-strange/mordo/helpers"
)

func Test_ProcessWatermark(t *testing.T) {
	testCases := []struct {
		name    string
		img     *[]byte
		xString string
		yString string
		scale   float64
		want    *[]byte
	}{
		{
			name:    "origin_2k.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			xString: "right",
			yString: "bottom",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_2k.jpg"),
		},
		{
			name:    "origin_2k_scale.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			xString: "right",
			yString: "bottom",
			scale:   0.5,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_2k_scale.jpg"),
		},
		{
			name:    "origin_2k_center.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			xString: "center",
			yString: "center",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_2k_center.jpg"),
		},
		{
			name:    "origin_4k.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			xString: "left",
			yString: "top",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_4k.jpg"),
		},
		{
			name:    "origin_big.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_big.jpg"),
			xString: "left",
			yString: "center",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_big.jpg"),
		},
		{
			name:    "origin_2k.png",
			img:     helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			xString: "right",
			yString: "bottom",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_2k.png"),
		},
		{
			name:    "origin_2k_scale.png",
			img:     helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			xString: "right",
			yString: "bottom",
			scale:   0.5,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_2k_scale.png"),
		},
		{
			name:    "origin_2k_center.png",
			img:     helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			xString: "center",
			yString: "center",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_2k_center.png"),
		},
		{
			name:    "origin_2k.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			xString: "right",
			yString: "bottom",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_2k.webp"),
		},
		{
			name:    "origin_2k_scale.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			xString: "right",
			yString: "bottom",
			scale:   0.5,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_2k_scale.webp"),
		},
		{
			name:    "origin_2k_center.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			xString: "center",
			yString: "center",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_2k_center.webp"),
		},
		{
			name:    "origin_4k.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			xString: "left",
			yString: "top",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_4k.webp"),
		},
		{
			name:    "origin_big.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_big.webp"),
			xString: "left",
			yString: "center",
			scale:   0,
			want:    helpers.ReadToBuff("../../tests/img/process/watermarked/origin_big.webp"),
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable https://blog.golang.org/subtests
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i, err := New(tc.img)
			if err != nil {
				t.Fatal(err)
			}

			// i.Resize("w", 700)
			// i.SCrop(700, 400)
			if err := i.Watermark(tc.xString, tc.yString, tc.scale); err != nil {
				t.Fatal(err)
			}

			if got, _, err := i.Process(); err != nil {
				t.Fatal(err)
			} else if !reflect.DeepEqual(got, tc.want) {
				t.Error("Watermarked image is not eq with test image")
				t.Error("Will write images to ../../tmp/ for further checks")
				p := fmt.Sprintf("../../tmp/%s", tc.name)
				ioutil.WriteFile(p, *got, 0644)
			}
		})

	}
}

func Benchmark_ProcessWatermark(b *testing.B) {
	testCases := []struct {
		name    string
		img     *[]byte
		xString string
		yString string
		scale   float64
	}{
		{
			name:    "origin_2k.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			xString: "right",
			yString: "bottom",
			scale:   0,
		},
		{
			name:    "origin_2k_scale.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			xString: "right",
			yString: "bottom",
			scale:   0.5,
		},
		{
			name:    "origin_2k_center.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			xString: "center",
			yString: "center",
			scale:   0,
		},
		{
			name:    "origin_4k.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			xString: "left",
			yString: "top",
			scale:   0,
		},
		{
			name:    "origin_big.jpg",
			img:     helpers.ReadToBuff("../../tests/i-jpg/origin_big.jpg"),
			xString: "left",
			yString: "center",
			scale:   0,
		},
		{
			name:    "origin_2k.png",
			img:     helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			xString: "right",
			yString: "bottom",
			scale:   0,
		},
		{
			name:    "origin_2k_scale.png",
			img:     helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			xString: "right",
			yString: "bottom",
			scale:   0.5,
		},
		{
			name:    "origin_2k_center.png",
			img:     helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			xString: "center",
			yString: "center",
			scale:   0,
		},
		{
			name:    "origin_2k.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			xString: "right",
			yString: "bottom",
			scale:   0,
		},
		{
			name:    "origin_2k_scale.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			xString: "right",
			yString: "bottom",
			scale:   0.5,
		},
		{
			name:    "origin_2k_center.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			xString: "center",
			yString: "center",
			scale:   0,
		},
		{
			name:    "origin_4k.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			xString: "left",
			yString: "top",
			scale:   0,
		},
		{
			name:    "origin_big.webp",
			img:     helpers.ReadToBuff("../../tests/i-webp/origin_big.webp"),
			xString: "left",
			yString: "center",
			scale:   0,
		},
	}

	b.ReportAllocs()
	b.SetBytes(2)
	b.ResetTimer()

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				i, err := New(tc.img)
				if err != nil {
					b.Fatal(err)
				}
				b.StartTimer()

				if err := i.Watermark(tc.xString, tc.yString, tc.scale); err != nil {
					b.Fatal(err)
				}

				if _, _, err = i.Process(); err != nil {
					b.Fatal(err)
				}
			}
		})

	}
}
