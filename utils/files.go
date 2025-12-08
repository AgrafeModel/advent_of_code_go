package utils

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func isTestEnv() bool {
	return flag.Lookup("test.v") != nil
}

func GetInputPath(year, day int) string {
	var r string
	if isTestEnv() {
		r = fmt.Sprintf("../../inputs/y%d/day%d.txt", year, day)
	} else {
		r = fmt.Sprintf("./inputs/y%d/day%d.txt", year, day)
	}
	fmt.Println(r)
	return r
}

func ReadFilePerLines(path string, fn func(line string)) {
	file, err := os.Open(path)
	HandleErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	for { // for each lines
		l, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		//remove the \n
		l = l[:len(l)-1]

		//show the line not formated to show \n etx..
		// fmt.Println(l)
		fn(l)
	}
}

// same as ReadFilePerLines but start from last line of the file
func ReadFilePerLinesReverse(path string, fn func(line string)) {
	file, err := os.Open(path)
	HandleErr(err)
	defer file.Close()

	stat, err := file.Stat()
	HandleErr(err)

	var size int64 = stat.Size()
	var offset int64 = 0
	var chunkSize int64 = 1024
	var buffer []byte
	var line string

	for offset < size {
		if size-offset < chunkSize {
			chunkSize = size - offset
		}
		offset += chunkSize
		buffer = make([]byte, chunkSize)
		_, err := file.ReadAt(buffer, size-offset)
		HandleErr(err)

		for i := len(buffer) - 1; i >= 0; i-- {
			if buffer[i] == '\n' {
				if len(line) > 0 {
					fn(reverseString(line))
					line = ""
				}
			} else {
				line += string(buffer[i])
			}
		}
	}
	if len(line) > 0 {
		fn(reverseString(line))
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
