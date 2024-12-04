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
