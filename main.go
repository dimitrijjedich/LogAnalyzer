package main

import (
	"bufio"
	"fmt"
	"os"
)

type Entry struct {
	time string
	log  string
}

func main() {
	readFile, err := os.Open("test.log")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	readFile.Close()
}
