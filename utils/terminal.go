package utils

import (
	"fmt"
	"os"
)

var DEBUG *bool

func IsDebugMode() bool {
	return DEBUG != nil && *DEBUG
}

func SetDebugMode(debug *bool) {
	DEBUG = debug
}

func Puzzle(str string) {
	fmt.Println("\033[32m [PUZZLE] " + str + "\033[0m")
}

func Step(str string, duration int) {
	fmt.Println("\033[34m [STEP] " + str + " ->  Duration: " + fmt.Sprintf("%d", duration) + "ms \033[0m")
}

func Debug(str any) {
	fmt.Println("\033[33m [DEBUG] ", str, "\033[0m")
}

func HandleErr(err error) {
	if err != nil {

		fmt.Println("\033[31m [ERROR] " + err.Error() + "\033[0m")
		os.Exit(0)
	}
}
