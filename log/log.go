package log

import "fmt"

var isVerbose bool = false

func SetVerbose(verbose bool) {
	isVerbose = verbose
}

func Log(msg string) {
	if isVerbose {
		println(msg)
	}
}

func Logf(format string, a ...interface{}) {
	if isVerbose {
		fmt.Printf(format, a...)
	}
}
