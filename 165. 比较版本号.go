package main

import "strings"
import "strconv"

func compareVersion(version1 string, version2 string) int {
	v1str := strings.Split(version1, ".")
	v2str := strings.Split(version2, ".")

	v1 := make([]int, 0, len(v1str))
	v2 := make([]int, 0, len(v2str))

	for i := 0; i < len(v1str); i++ {
		num, _ := strconv.Atoi(v1str[i])
		v1 = append(v1, num)
	}
	for i := 0; i < len(v2str); i++ {
		num, _ := strconv.Atoi(v2str[i])
		v2 = append(v2, num)
	}

    length:=0
	if len(v1) > len(v2) {
		length = len(v2)
	} else {
		length = len(v1)
	}

	for i := 0; i < length; i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}

	if len(v1) > len(v2) {
		for i := length; i < len(v1); i++ {
			if v1[i] > 0 {
				return 1
			}
		}
	} else {
		for i := length; i < len(v2); i++ {
			if v2[i] > 0 {
				return -1
			}
		}
	}

	return 0
}