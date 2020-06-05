package cologger

import (
	"fmt"
	"log"
)

func out(color int, label string, v ...interface{}) {
	prefix := fmt.Sprintf("\x1b[%dm[%s]\x1b[%dm", color, label, reset)
	data := append([]interface{}{prefix}, v...)
	log.Println(data...)
}

func outf(color int, label string, format string, v ...interface{}) {
	prefix := fmt.Sprintf("\x1b[%dm[%s]\x1b[%dm", color, label, reset)
	body := fmt.Sprintf(format, v...)
	log.Println(prefix, body)
}
