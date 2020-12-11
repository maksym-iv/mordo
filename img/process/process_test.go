// +build !test

package process

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/xmackex/mordo/helpers"
)

func Benchmark_New(b *testing.B) {
	testCases := []struct {
		name   string
		imgBuf []byte
	}{
		{
			name:   "origin_2k",
			imgBuf: helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
		},
	}

	b.ReportAllocs()
	b.SetBytes(2)
	b.ResetTimer()

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				New(tc.imgBuf)
			}
		})

	}
}

func Test_Process(t *testing.T) {
	testCases := []struct {
		name string
		img  []byte
		want []byte
	}{
		{
			name: "origin_2k.jpg",
			img:  helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
			want: helpers.ReadToBuff("../../tests/img/process/processed/origin_2k.jpg"),
		},
		{
			name: "origin_4k.jpg",
			img:  helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
			want: helpers.ReadToBuff("../../tests/img/process/processed/origin_4k.jpg"),
		},
		{
			name: "origin_big.jpg",
			img:  helpers.ReadToBuff("../../tests/i-jpg/origin_big.jpg"),
			want: helpers.ReadToBuff("../../tests/img/process/processed/origin_big.jpg"),
		},
		{
			name: "origin_2k.png",
			img:  helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
			want: helpers.ReadToBuff("../../tests/img/process/processed/origin_2k.png"),
		},
		{
			name: "origin_2k.webp",
			img:  helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
			want: helpers.ReadToBuff("../../tests/img/process/processed/origin_2k.webp"),
		},
		{
			name: "origin_4k.webp",
			img:  helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
			want: helpers.ReadToBuff("../../tests/img/process/processed/origin_4k.webp"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i, err := New(tc.img)
			if err != nil {
				t.Fatal(err)
			}

			if got, _, err := i.Process(); err != nil {
				t.Fatal(err)
			} else if !reflect.DeepEqual(got, tc.want) {
				t.Error("Processed image is not eq with test image")
				t.Error("Will write images to ../../tmp/ for further checks")
				p := fmt.Sprintf("../../tmp/%s", tc.name)
				ioutil.WriteFile(p, got, 0644)
			}
		})

	}
}

func Benchmark_Process(b *testing.B) {
	testCases := []struct {
		name string
		img  []byte
	}{
		{
			name: "origin_2k.jpg",
			img:  helpers.ReadToBuff("../../tests/i-jpg/origin_2k.jpg"),
		},
		{
			name: "origin_4k.jpg",
			img:  helpers.ReadToBuff("../../tests/i-jpg/origin_4k.jpg"),
		},
		{
			name: "origin_big.jpg",
			img:  helpers.ReadToBuff("../../tests/i-jpg/origin_big.jpg"),
		},
		{
			name: "origin_2k.png",
			img:  helpers.ReadToBuff("../../tests/i-png/origin_2k.png"),
		},
		{
			name: "origin_2k.webp",
			img:  helpers.ReadToBuff("../../tests/i-webp/origin_2k.webp"),
		},
		{
			name: "origin_4k.webp",
			img:  helpers.ReadToBuff("../../tests/i-webp/origin_4k.webp"),
		},
	}

	b.ReportAllocs()
	b.SetBytes(2)
	b.ResetTimer()

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// b.StopTimer()
				i, err := New(tc.img)
				if err != nil {
					b.Fatal(err)
				}
				// b.StartTimer()

				_, _, err = i.Process()
				if err != nil {
					b.Fatal(err)
				}
			}
		})

	}
}
