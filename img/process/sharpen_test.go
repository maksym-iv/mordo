// +build !test

package process

import (
	"fmt"
	"io/ioutil"
	// "fmt"
	// "io/ioutil"
	"reflect"
	"testing"

	"github.com/xmackex/mordo/helpers"
)

func Test_ProcessSharpen(t *testing.T) {
	testCases := []struct {
		name   string
		img    []byte
		radius float64
		X1     float64
		Y2     float64
		Y3     float64
		M1     float64
		M2     float64
		want   []byte
	}{
		{
			name:   "origin_2k.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
			want:   helpers.ReadToBuff("../../tests/img/process/sharpened/origin_2k.jpg"),
		},
		{
			name:   "origin_4k.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
			want:   helpers.ReadToBuff("../../tests/img/process/sharpened/origin_4k.jpg"),
		},
		{
			name:   "origin_big.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_big.jpg"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
			want:   helpers.ReadToBuff("../../tests/img/process/sharpened/origin_big.jpg"),
		},
		{
			name:   "origin_2k.png",
			img:    helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
			want:   helpers.ReadToBuff("../../tests/img/process/sharpened/origin_2k.png"),
		},
		{
			name:   "origin_2k.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
			want:   helpers.ReadToBuff("../../tests/img/process/sharpened/origin_2k.webp"),
		},
		{
			name:   "origin_4k.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
			want:   helpers.ReadToBuff("../../tests/img/process/sharpened/origin_4k.webp"),
		},
		{
			name:   "origin_big.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_big.webp"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
			want:   helpers.ReadToBuff("../../tests/img/process/sharpened/origin_big.webp"),
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

			if err := i.Sharpen(tc.radius, tc.X1, tc.Y2, tc.Y3, tc.M1, tc.M2); err != nil {
				t.Fatal(err)
			}

			if got, _, err := i.Process(); err != nil {
				t.Fatal(err)
			} else if !reflect.DeepEqual(got, tc.want) {
				t.Error("Resized image is not eq with test image")
				t.Error("Will write images to ../../tmp/ for further checks")
				p := fmt.Sprintf("../../tmp/%s", tc.name)
				ioutil.WriteFile(p, got, 0644)
			}
		})

	}
}

func Benchmark_ProcessSharpen(b *testing.B) {
	testCases := []struct {
		name   string
		img    []byte
		radius float64
		X1     float64
		Y2     float64
		Y3     float64
		M1     float64
		M2     float64
	}{
		{
			name:   "origin_2k.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
		},
		{
			name:   "origin_4k.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
		},
		{
			name:   "origin_big.jpg",
			img:    helpers.ReadToBuff("../../tests/i-jpg/origin_big.jpg"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
		},
		{
			name:   "origin_2k.png",
			img:    helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
		},
		{
			name:   "origin_2k.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
		},
		{
			name:   "origin_4k.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
		},
		{
			name:   "origin_big.webp",
			img:    helpers.ReadToBuff("../../tests/i-webp/origin_big.webp"),
			radius: 1,
			X1:     2,
			Y2:     1,
			Y3:     5,
			M1:     0,
			M2:     3,
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

				if err := i.Sharpen(tc.radius, tc.X1, tc.Y2, tc.Y3, tc.M1, tc.M2); err != nil {
					b.Fatal(err)
				}

				if _, _, err = i.Process(); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
