package cologger

import "os"

func Info(v ...interface{}) {
	out(greenBack, "INFO", v...)
}

func Infof(format string, v ...interface{}) {
	outf(greenBack, "INFO", format, v...)
}

func Warn(v ...interface{}) {
	out(yellowBack, "WARN", v...)
}

func Warnf(format string, v ...interface{}) {
	outf(yellowBack, "WARN", format, v...)
}

func Fatal(v ...interface{}) {
	out(redBack, "FATAL", v...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	outf(redBack, "FATAL", format, v...)
	os.Exit(1)
}
