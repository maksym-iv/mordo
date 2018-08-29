package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_process(t *testing.T) {
	testCases := []struct {
		name      string
		image     string
		qs        map[string]string
		stageVars map[string]string
		want      []byte
	}{
		{
			name:  "origin_2k.jpg",
			image: "origin_2k.jpg",
			// image: "tests/read/origin_2k.jpg",
			qs: map[string]string{
				"width": "700",
				"w_x":   "right",
				"w_y":   "bottom",
				// "w_s":   "0.9",
				// "sc_x": "400",
				// "sc_y": "400",
			},
			stageVars: map[string]string{
				"s3_bucket":        "demo-mordo-images",
				"s3_bucket_region": "us-west-1",
			},
			// want:    helpers.ReadToBuff("../../tests/io/s3/read/origin_2k.jpg"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if r, err := process(tc.image, tc.qs, tc.stageVars); err != nil {
				t.Fatal(err)
			} else {
				b, err := base64.StdEncoding.DecodeString(r.Body)
				if err != nil {
					t.Fatal(err)
				}
				p := fmt.Sprintf("../../tmp/%s", tc.name)
				ioutil.WriteFile(p, b, 0644)
			}
		})
	}
}

func Benchmark_process(b *testing.B) {
	testCases := []struct {
		name      string
		image     string
		qs        map[string]string
		stageVars map[string]string
	}{
		{
			name:  "origin_light_2k.jpg",
			image: "origin_light_2k.jpg",
			// image: "tests/read/origin_2k.jpg",
			qs: map[string]string{
				"width": "1000",
				// "w_x": "right",
				// "w_y": "bottom",
				// "w_s":   "0.9",
				// "sc_x": "400",
				// "sc_y": "400",
			},
			stageVars: map[string]string{
				"s3_bucket":        "dev-mordo-images",
				"s3_bucket_region": "us-west-1",
			},
		},
	}

	b.ReportAllocs()
	b.SetBytes(2)
	b.ResetTimer()

	for _, tc := range testCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if _, err := process(tc.image, tc.qs, tc.stageVars); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
