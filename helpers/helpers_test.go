package helpers

import (
	"errors"
	"reflect"
	"testing"
)

func Test_FormatBodyJSON(t *testing.T) {
	testCases := []struct {
		name string
		msg  string
		err  error
		want string
	}{
		{
			name: "1",
			err:  errors.New("[] json: unsupported type: map[int]main.Foo"),
			want: `{"error":"[] json: unsupported type: map[int]main.Foo"}`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := FormatBodyJSON(tc.err); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%v NE %v", got, tc.want)
			}
		})
	}
}

func Test_ContainsString(t *testing.T) {
	testCases := []struct {
		name string
		elem string
		arr  []string
		want bool
	}{
		{
			name: "1",
			arr:  []string{"q", "w", "e"},
			elem: "q",
			want: true,
		},
		{
			name: "2",
			arr:  []string{"q", "w", "e"},
			elem: "rr",
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := ContainsString(tc.arr, tc.elem); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%v NE %v", got, tc.want)
			}
		})
	}
}

func Test_ContainsInt(t *testing.T) {
	testCases := []struct {
		name string
		elem int
		arr  []int
		want bool
	}{
		{
			name: "1",
			arr:  []int{1, 2, 3},
			elem: 1,
			want: true,
		},
		{
			name: "2",
			arr:  []int{1, 2, 3},
			elem: 22,
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := ContainsInt(tc.arr, tc.elem); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%v NE %v", got, tc.want)
			}
		})
	}
}
