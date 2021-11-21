package tools

import "sort"

func IsKeyInSlice(s []string, key string) bool {
	for _, i := range s {
		if key == i {
			return true
		}
	}
	return false
}

func IsIntKeyInSlice(s []int, key int) bool {
	for _, i := range s {
		if key == i {
			return true
		}
	}
	return false
}

func GussMinNum(i []int) int {
	var tmpSlice = make([]int, len(i))
	copy(tmpSlice, i)
	sort.Ints(tmpSlice)
	maxNum := tmpSlice[len(tmpSlice)-1]
	if maxNum == len(tmpSlice) {
		return maxNum + 1
	}
	for i := 1; i <= maxNum; i++ {
		if !IsIntKeyInSlice(tmpSlice, i) {
			return i
		}
	}
	return 1
}

func UniqStringSlice(s []string) []string {
	d := make([]string, 0)
	tempMap := make(map[string]bool, len(s))
	for _, v := range s { // 以值作为键名
		if tempMap[v] == false {
			tempMap[v] = true
			d = append(d, v)
		}
	}
	return d
}

func DeleteObjInTargetSlice(s, t []string) {
	for i := 0; i < len(s); i++ {
		if IsKeyInSlice(t, s[i]) {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
}
