package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	time string
	log  string
}

func main() {
	var logs []Entry
	readFile, err := os.Open("test.log")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		var line = strings.SplitN(fileScanner.Text(), "] ", 2)
		var entry = Entry{
			time: strings.Trim(line[0], "["),
			log:  line[1]}
		logs = append(logs, entry)
	}
	readFile.Close()
}
