package utils

import (
	"log"
	"fmt"
	"testing"
)

func FatalOnErr(err error, format string, v ...interface{}) bool {
	if err != nil {
		log.Fatalf("%s: %s \n", fmt.Sprintf(format, v...), err)
		return true
	}
	return false
}

func PanicOnErr(err error, format string, v ...interface{}) bool {
	if err != nil {
		log.Panicf(fmt.Sprintf("%s: %s \n", fmt.Sprintf(format, v...), err))
		return true
	}
	return false
}

func LogOnErr(err error, format string, v ...interface{}) bool {
	if err != nil {
		log.Printf("%s: %s \n", fmt.Sprintf(format, v...), err)
		return true
	}
	return false
}

func PrintOnErr(err error, format string, v ...interface{}) bool {
	if err != nil {
		fmt.Printf("%s: %s \n", fmt.Sprintf(format, v...), err)
		return true
	}
	return false
}

func FailTestOnErr(err error, tb testing.TB, format string, v ...interface{}) {
	if err != nil {
		tb.Errorf("%s: %s \n", fmt.Sprintf(format, v...), err)
	}
}