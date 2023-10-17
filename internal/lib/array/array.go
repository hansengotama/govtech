package array

import "strings"

func IsContain(arr []string, target string) bool {
	for _, s := range arr {
		if strings.ToLower(s) == strings.ToLower(target) {
			return true
		}
	}

	return false
}
