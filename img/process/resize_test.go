// +build !test

package process

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"gitlab.com/doctor-strange/mordo/helpers"
)

func Test_ProcessResize(t *testing.T) {
	testCases := []struct {
		name      string
		img       *[]byte
		dimension string
		size      int
		want      *[]byte
	}{
		{
			name:      "origin_2k_width.jpg",
			img:       helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			dimension: "width",
			size:      500,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_2k_width.jpg"),
		},
		{
			name:      "origin_2k_height.jpg",
			img:       helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			dimension: "height",
			size:      500,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_2k_height.jpg"),
		},
		{
			name:      "origin_4k_h.jpg",
			img:       helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			dimension: "h",
			size:      1000,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_4k_h.jpg"),
		},
		{
			name:      "origin_big_w.jpg",
			img:       helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			dimension: "w",
			size:      1000,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_big_w.jpg"),
		},
		{
			name:      "origin_2k_width.png",
			img:       helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			dimension: "width",
			size:      500,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_2k_width.png"),
		},
		{
			name:      "origin_2k_height.png",
			img:       helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			dimension: "height",
			size:      500,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_2k_height.png"),
		},
		{
			name:      "origin_2k_width.webp",
			img:       helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			dimension: "width",
			size:      500,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_2k_width.webp"),
		},
		{
			name:      "origin_2k_height.webp",
			img:       helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			dimension: "height",
			size:      500,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_2k_height.webp"),
		},
		{
			name:      "origin_4k_h.webp",
			img:       helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			dimension: "h",
			size:      1000,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_4k_h.webp"),
		},
		{
			name:      "origin_big_w.webp",
			img:       helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			dimension: "w",
			size:      1000,
			want:      helpers.ReadToBuff("../../tests/img/process/resized/origin_big_w.webp"),
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

			if err := i.Resize(tc.dimension, tc.size); err != nil {
				t.Fatal(err)
			}

			if got, _, err := i.Process(); err != nil {
				t.Fatal(err)
			} else if !reflect.DeepEqual(got, tc.want) {
				t.Error("Resized image is not eq with test image")
				t.Error("Will write images to ../../tmp/ for further checks")
				p := fmt.Sprintf("../../tmp/%s", tc.name)
				ioutil.WriteFile(p, *got, 0644)
			}
		})
	}
}

func Benchmark_ProcessResize(b *testing.B) {
	testCases := []struct {
		name      string
		img       *[]byte
		dimension string
		size      int
	}{
		{
			name:      "origin_2k_width.jpg",
			img:       helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			dimension: "width",
			size:      500,
		},
		{
			name:      "origin_4k_h.jpg",
			img:       helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			dimension: "h",
			size:      1000,
		},
		{
			name:      "origin_big_w.jpg",
			img:       helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			dimension: "w",
			size:      1000,
		},
		{
			name:      "origin_2k_width.png",
			img:       helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			dimension: "width",
			size:      500,
		},
		{
			name:      "origin_2k_width.webp",
			img:       helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			dimension: "width",
			size:      500,
		},
		{
			name:      "origin_4k_h.webp",
			img:       helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			dimension: "h",
			size:      1000,
		},
		{
			name:      "origin_big_w.webp",
			img:       helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			dimension: "w",
			size:      1000,
		},
	}

	b.ReportAllocs()
	b.SetBytes(2)
	b.ResetTimer()

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				i, err := New(tc.img)
				if err != nil {
					b.Fatal(err)
				}

				if err := i.Resize(tc.dimension, tc.size); err != nil {
					b.Fatal(err)
				}

				if _, _, err = i.Process(); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
