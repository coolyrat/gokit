package utils

import "strconv"

// 字符串slice转int slice
func StrSliceToIntSlice(s []string) []int {
	if s == nil || len(s) == 0 {
		return make([]int, 0)
	}

	rst := make([]int, 0, len(s))
	for _, v := range s {
		val, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		rst = append(rst, val)
	}
	return rst
}