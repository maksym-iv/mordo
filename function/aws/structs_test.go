package main

import (
	"reflect"
	"testing"
)

func Test_newQs(t *testing.T) {
	testCases := []struct {
		name  string
		image string
		qs    map[string]string
		want  *QS
	}{
		{
			name:  "origin_2k.jpg",
			image: "origin_2k.jpg",
			qs: map[string]string{
				"width":   "101",
				"w_x":     "right",
				"w_y":     "bottom",
				"dpr":     "2",
				"sharpen": "t",
				"sc_x":    "600",
				"sc_y":    "300",
				"c_left":  "300",
				"c_top":   "400",
				"c_x":     "200",
				"c_y":     "600",
			},
			want: &QS{
				Image:     "origin_2k.jpg",
				Resize:    Resize{Width: 101},
				Watermark: Watermark{WX: "right", WY: "bottom"},
				SCrop:     SCrop{Width: 600, Heigth: 300},
				Crop:      Crop{Left: 300, Top: 400, Width: 200, Heigth: 600},
				DPR:       2,
				Sharpen:   true,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got, err := newQs(tc.image, tc.qs); err != nil {
				t.Fatal(err)
			} else if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got: %+v, Want: %+v", got, tc.want)
			} else {
				// t.Logf("%+v", got)
			}

		})
	}
}

func Test_hashQss(t *testing.T) {
	testCases := []struct {
		name string
		qss  *QS
		want string
	}{
		{
			name: "origin_2k.jpg 101",
			qss: &QS{
				Image:     "origin_2k.jpg",
				Resize:    Resize{Width: 101},
				Watermark: Watermark{WX: "right", WY: "bottom"},
				DPR:       2,
				Sharpen:   true,
			},
			want: "prc_5112ac2e71dcfe5de42c580b0a01089106231d53_origin_2k.jpg",
		},
		{
			name: "origin_2k.jpg 200",
			qss: &QS{
				Image:     "origin_2k.jpg",
				Resize:    Resize{Width: 200},
				Watermark: Watermark{WX: "right", WY: "bottom"},
				DPR:       2,
				Sharpen:   true,
			},
			want: "prc_f98fe55fc7400ab17577c16e54544e09c884931a_origin_2k.jpg",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// got := tc.qss.hashQss()
			// t.Log(got)
			if got := tc.qss.hashPath(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Sha1 hashes differ. Got: %s, Want: %s", got, tc.want)
			}
		})
	}
}

func Benchmark_newPath(b *testing.B) {
	testCases := []struct {
		name string
		qss  *QS
		want string
	}{
		{
			name: "origin_2k.jpg",
			qss: &QS{
				Image:     "origin_2k.jpg",
				Resize:    Resize{Width: 101},
				Watermark: Watermark{WX: "right", WY: "bottom"},
				DPR:       2,
				Sharpen:   true,
			},
			// want: &QS{
			// 	Image:     "origin_2k.jpg",
			// 	Resize:    Resize{Width: 101},
			// 	Watermark: Watermark{WX: "right", WY: "bottom"},
			// 	DPR:       2,
			// 	Sharpen:   true,
			// },
		},
	}

	for _, tc := range testCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.qss.hashPath()
			}
		})
	}
}

func Test_Error(t *testing.T) {
	testCases := []struct {
		name string
		err  *gwError
		want string
	}{
		{
			name: "1",
			err:  newErr("something went wrong"),
			want: "something went wrong",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := tc.err.Error(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Error text differs. Got: %s, Want: %s", got, tc.want)
			}
		})
	}
}

func Test_ErrorJSON(t *testing.T) {
	testCases := []struct {
		name string
		err  *gwError
		want string
	}{
		{
			name: "1",
			err:  newErr("something went wrong"),
			want: `{"err":"something went wrong"}`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := tc.err.ErrorJSON(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Error text differs. Got: %s, Want: %s", got, tc.want)
			}
		})
	}
}
