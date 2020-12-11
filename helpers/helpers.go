package helpers

import (
	"io/ioutil"
)

func ReadToBuff(p string) []byte {
	b, _ := ioutil.ReadFile(p)
	// TODO: uncomment
	// if err != nil {
	// 	panic(err)
	// }

	return b
}

func ContainsString(arr []string, elem string) bool {
	for i := range arr {
		if arr[i] == elem {
			return true
		}
	}

	return false
}

func ContainsInt(arr []int, elem int) bool {
	for i := range arr {
		if arr[i] == elem {
			return true
		}
	}

	return false
}
