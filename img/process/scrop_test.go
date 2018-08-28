// +build !test

package process

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"gitlab.com/doctor-strange/mordo/helpers"
)

func Test_ProcessSCrop(t *testing.T) {
	testCases := []struct {
		name          string
		img           *[]byte
		width, height int
		want          *[]byte
	}{
		{
			name:   "origin_light_2k.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_light_2k.jpg"),
			width:  1000,
			height: 500,
			want:   helpers.ReadToBuff("../../tests/img/process/scropped/origin_light_2k.jpg"),
		},
		{
			name:   "origin_4k.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			width:  600,
			height: 300,
			want:   helpers.ReadToBuff("../../tests/img/process/scropped/origin_4k.jpg"),
		},
		{
			name:   "origin_light_big.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_light_big.jpg"),
			width:  600,
			height: 300,
			want:   helpers.ReadToBuff("../../tests/img/process/scropped/origin_light_big.jpg"),
		},
		{
			name:   "origin_light_2k.png",
			img:    helpers.ReadToBuff("../../tests/i-png/origin_light_2k.png"),
			width:  1000,
			height: 500,
			want:   helpers.ReadToBuff("../../tests/img/process/scropped/origin_light_2k.png"),
		},
		{
			name:   "origin_light_2k.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_light_2k.webp"),
			width:  1000,
			height: 500,
			want:   helpers.ReadToBuff("../../tests/img/process/scropped/origin_light_2k.webp"),
		},
		{
			name:   "origin_4k.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			width:  600,
			height: 300,
			want:   helpers.ReadToBuff("../../tests/img/process/scropped/origin_4k.webp"),
		},
		{
			name:   "origin_light_big.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_light_big.webp"),
			width:  600,
			height: 300,
			want:   helpers.ReadToBuff("../../tests/img/process/scropped/origin_light_big.webp"),
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

			// i.Resize("w", 1000)
			if err := i.SCrop(tc.width, tc.height); err != nil {
				t.Fatal(err)
			}

			if got, _, err := i.Process(); err != nil {
				t.Fatal(err)
			} else if !reflect.DeepEqual(got, tc.want) {
				t.Error("SCropped image is not eq with test image")
				t.Error("Will write images to ../../tmp/ for further checks")
				p := fmt.Sprintf("../../tmp/%s", tc.name)
				ioutil.WriteFile(p, *got, 0644)
			}
		})

	}
}

func Benchmark_ProcessSCrop(b *testing.B) {
	testCases := []struct {
		name          string
		img           *[]byte
		width, height int
	}{
		{
			name:   "origin_light_2k.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_light_2k.jpg"),
			width:  1000,
			height: 500,
		},
		{
			name:   "origin_4k.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			width:  600,
			height: 300,
		},
		{
			name:   "origin_light_big.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_light_big.jpg"),
			width:  600,
			height: 300,
		},
		{
			name:   "origin_light_2k.png",
			img:    helpers.ReadToBuff("../../tests/i-png/origin_light_2k.png"),
			width:  1000,
			height: 500,
		},
		{
			name:   "origin_light_2k.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_light_2k.webp"),
			width:  1000,
			height: 500,
		},
		{
			name:   "origin_4k.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			width:  600,
			height: 300,
		},
		{
			name:   "origin_light_big.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_light_big.webp"),
			width:  600,
			height: 300,
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

				if err := i.SCrop(tc.width, tc.height); err != nil {
					b.Fatal(err)
				}

				if _, _, err = i.Process(); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
