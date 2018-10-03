package s3

import (
	"fmt"
	"io/ioutil"
	"reflect"
	// "os"
	// "strconv"
	"testing"
	// "github.com/xmackex/mordo/base"

	"github.com/xmackex/mordo/helpers"
)

func Benchmark_NewImg(b *testing.B) {
	testCases := []struct {
		name    string
		bucket  string
		imgPath string
	}{
		{
			name:    "origin_2k.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/origin_2k.jpg",
		},
		// {
		// 	name:    "origin_4k.jpg",
		// 	bucket:  "dev-mordo-images",
		// 	imgPath: "tests/read/origin_4k.jpg",
		// 	want:    helpers.ReadToBuff("../../tests/io/s3/read/origin_4k.jpg"),
		// },
		// {
		// 	name:    "origin_big.jpg",
		// 	bucket:  "dev-mordo-images",
		// 	imgPath: "tests/read/origin_big.jpg",
		// 	want:    helpers.ReadToBuff("../../tests/io/s3/read/origin_big.jpg"),
		// },
	}

	for _, tc := range testCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := NewImg(tc.bucket, tc.imgPath, "us-west-1")
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func Test_Read(t *testing.T) {
	testCases := []struct {
		name    string
		bucket  string
		imgPath string
		want    []byte
	}{
		{
			name:    "origin_2k.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/origin_2k.jpg",
			want:    helpers.ReadToBuff("../../tests/io/s3/read/origin_2k.jpg"),
		},
		// {
		// 	name:    "origin_4k.jpg",
		// 	bucket:  "dev-mordo-images",
		// 	imgPath: "tests/read/origin_4k.jpg",
		// 	want:    helpers.ReadToBuff("../../tests/io/s3/read/origin_4k.jpg"),
		// },
		// {
		// 	name:    "origin_big.jpg",
		// 	bucket:  "dev-mordo-images",
		// 	imgPath: "tests/read/origin_big.jpg",
		// 	want:    helpers.ReadToBuff("../../tests/io/s3/read/origin_big.jpg"),
		// },
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i, err := NewImg(tc.bucket, tc.imgPath, "us-west-1")
			if err != nil {
				t.Fatal(err)
			}
			if err := i.Read(); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(i.Buff, tc.want) {
				t.Error("Read image is not eq with test image")
				t.Error("Will write images to ../../tmp/ for further checks")
				p := fmt.Sprintf("../../tmp/%s", tc.name)
				ioutil.WriteFile(p, *i.Buff, 0644)
			}
		})
	}
}

func Benchmark_Read(b *testing.B) {
	testCases := []struct {
		name    string
		bucket  string
		imgPath string
	}{
		{
			name:    "origin_2k.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/origin_2k.jpg",
		},
	}

	b.ReportAllocs()
	b.SetBytes(2)
	b.ResetTimer()

	for _, tc := range testCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				i, err := NewImg(tc.bucket, tc.imgPath, "us-west-1")
				if err != nil {
					b.Fatal(err)
				}
				if err := i.Read(); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func Test_IsExist(t *testing.T) {
	testCases := []struct {
		name    string
		bucket  string
		imgPath string
		want    bool
	}{
		{
			name:    "Exist origin_2k.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "origin_2k.jpg",
			want:    true,
		},
		{
			name:    "Not exist 43.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/43.jpg",
			want:    false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i, err := NewImg(tc.bucket, tc.imgPath, "us-west-1")
			if err != nil {
				t.Fatal(err)
			}
			if got, err := i.IsExist(tc.imgPath); err != nil {
				t.Fatal(err)
			} else if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Exist error. Got: %t, Want: %t", got, tc.want)
			}
		})
	}
}

func Benchmark_Exists(b *testing.B) {
	testCases := []struct {
		name    string
		bucket  string
		imgPath string
	}{
		{
			name:    "Exist origin_2k.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/origin_2k.jpg",
		},
		{
			name:    "Exist origin_4k.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/origin_4k.jpg",
		},
		{
			name:    "Exist origin_big.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/origin_big.jpg",
		},
		{
			name:    "Not exist 43_2k.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/43_2k.jpg",
		},
		{
			name:    "Not exist 43_4k.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/43_4k.jpg",
		},
		{
			name:    "Not exist 43_big.jpg",
			bucket:  "dev-mordo-images",
			imgPath: "tests/read/43_big.jpg",
		},
	}

	b.ReportAllocs()
	b.SetBytes(2)
	b.ResetTimer()

	for _, tc := range testCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				i, err := NewImg(tc.bucket, tc.imgPath, "us-west-1")
				if err != nil {
					b.Fatal(err)
				}
				if _, err := i.IsExist(tc.imgPath); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func Test_Write(t *testing.T) {
	testCases := []struct {
		name    string
		bucket  string
		imgBuff []byte
		imgPath string
		want    []byte
	}{
		{
			name:    "origin_2k.jpg",
			bucket:  "dev-mordo-images",
			imgBuff: helpers.ReadToBuff("../../tests/origin_2k.jpg"),
			imgPath: "tests/wrote/origin_2k.jpg",
			want:    helpers.ReadToBuff("../../tests/io/s3/wrote/origin_2k.jpg"),
		},
		// {
		// 	name:    "origin_4k.jpg",
		// 	bucket:  "dev-mordo-images",
		// 	imgBuff: helpers.ReadToBuff("../../tests/origin_4k.jpg"),
		// 	imgPath: "tests/wrote/origin_4k.jpg",
		// 	want:    helpers.ReadToBuff("../../tests/io/s3/wrote/origin_4k.jpg"),
		// },
		// {
		// 	name:    "origin_big.jpg",
		// 	bucket:  "dev-mordo-images",
		// 	imgBuff: helpers.ReadToBuff("../../tests/origin_big.jpg"),
		// 	imgPath: "tests/wrote/origin_4k.jpg",
		// 	want:    helpers.ReadToBuff("../../tests/io/s3/wrote/origin_big.jpg"),
		// },
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i, err := NewImg(tc.bucket, tc.imgPath, "us-west-1")
			if err != nil {
				t.Fatal(err)
			}
			i.Img.UpdateBuff(tc.imgBuff)
			if err := i.Write(); err != nil {
				t.Fatal(err)
			}
			if err := i.Read(); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(i.Buff, tc.want) {
				t.Error("Read image is not eq with test image")
				t.Error("Will write images to ../../tmp/ for further checks")
				p := fmt.Sprintf("../../tmp/%s", tc.name)
				ioutil.WriteFile(p, *i.Buff, 0644)
			}
		})
	}
}
