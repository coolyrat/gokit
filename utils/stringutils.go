package utils

import (
	"strconv"
)

// 字符串转int，失败使用默认值
func StrToIntDefault(s string, def int) int {
	if s == "" {
		return def
	}

	v, err := strconv.Atoi(s)
	if LogOnErr(err, "fail to convert %s to int value", s) {
		return def
	} else {
		return v
	}

}

// 字符串转int
func StrToInt(s string) int {
	return StrToIntDefault(s, 0)
}

// 字符串转bool，失败使用默认值
func StrToBoolDefault(s string, def bool) bool {
	if s == "" {
		return def
	}

	v, err := strconv.ParseBool(s)
	if LogOnErr(err, "fail to convert %s to bool value", s) {
		return def
	} else {
		return v
	}

}

// 字符串转bool
func StrToBool(s string) bool {
	return StrToBoolDefault(s, false)
}