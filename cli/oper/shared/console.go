package shared

import "github.com/KevinZonda/GoX/pkg/console"

func WarnF(format string, args ...interface{}) {
	console.Yellow.AsForeground().WriteLine(format, args...)
}

func ErrorF(format string, args ...interface{}) {
	console.Red.AsForeground().WriteLine(format, args...)
}

func SuccessF(format string, args ...interface{}) {
	console.Green.AsForeground().WriteLine(format, args...)
}

func InfoF(format string, args ...interface{}) {
	console.Blue.AsForeground().WriteLine(format, args...)
}

func HighlightF(format string, args ...interface{}) {
	console.Cyan.AsForeground().WriteLine(format, args...)
}
