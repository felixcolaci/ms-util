package logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func Info(logLine string) {
	line := ":: INFO :: " + logLine
	_, file, no, ok := runtime.Caller(1)
	go func() {

		if ok {
			line += " :: " + fmt.Sprintf("%v#%v", trimCaller(file), no)
		}
		log.Println(line)
	}()
}
func Warn(logLine string) {
	line := ":: WARN :: " + logLine
	_, file, no, ok := runtime.Caller(1)
	go func() {

		if ok {
			line += " :: " + fmt.Sprintf("%v#%v", trimCaller(file), no)
		}
		log.Println(line)
	}()
}
func Boot(logLine string) {
	line := ":: BOOT :: " + logLine
	_, file, no, ok := runtime.Caller(1)
	go func() {

		if ok {
			line += " :: " + fmt.Sprintf("%v#%v", trimCaller(file), no)
		}
		log.Println(line)
	}()
}
func Err(logLine string) {

	line := ":: ERRO :: " + logLine
	_, file, no, ok := runtime.Caller(1)
	go func() {
		if ok {
			line += " :: " + fmt.Sprintf("%v#%v", trimCaller(file), no)
		}
		log.Println(line)
	}()
}

func trimCaller(in string) string {

	slices := strings.Split(in, "/")

	return slices[len(slices)-1]
}
