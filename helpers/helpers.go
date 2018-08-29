package helpers

import (
	"io/ioutil"
	// "os"
)

// Transform - implements Query String (map[string]string) transformation to other formats

// GetEnv - get env key or default
// func GetEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}
// 	return fallback
// }

func ReadToBuff(p string) []byte {
	b, _ := ioutil.ReadFile(p)
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
