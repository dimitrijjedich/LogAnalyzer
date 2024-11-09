package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Entry struct {
	time string
	log  string
}

func readFile(filename string) ([]string, error) {
	var lines []string
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	regExp := regexp.MustCompile(`^\[\d{4}-\d{2}-\d{2}`)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if regExp.MatchString(line) {
			lines = append(lines, line)
		}
	}
	if err := fileScanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func parseEntries(lines []string) []Entry {
	var entries []Entry
	for _, line := range lines {
		var parts = strings.SplitN(line, "] ", 2)
		var entry = Entry{
			time: parts[0][12:],
			log:  parts[1]}
		entries = append(entries, entry)
	}
	return entries
}

func main() {
	lines, err := readFile("./test.log")
	if err != nil {
		fmt.Println("Error parsing file", err)
	}
	entries := parseEntries(lines)
	for _, entry := range entries {
		fmt.Printf("Time: %v -> %v \n\n", entry.time, entry.log)
	}
}
