package main

import (
	"bufio"
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

func Search(needle []rune, hayStack []rune) int {
	n := len(needle)
	maxLen := len(hayStack)
	skipTable := shiftArray(needle)

	position := n - 1
	for position < maxLen {
		println("Current position", position)
		i := n - 1
		for ; i >= 0 && needle[i] == hayStack[position]; i, position = i-1, position-1 {
			println("Inner for loop with i: ", i, " and postion: ", position)
		}
		if i == -1 {
			return position + 1
		}
		_, exist := skipTable[hayStack[position]]
		if exist {
			position = position + skipTable[hayStack[position]]
		} else {
			position = position + n
		}
	}
	return -1
}

func shiftArray(needle []rune) map[rune]int {
	result := make(map[rune]int)
	n := len(needle)
	for i := n - 1; i >= 0; i-- {
		_, exists := result[needle[i]]
		if !exists {
			result[needle[i]] = n - 1 - i
		}
	}
	return result
}

func main() {
	/*lines, err := readFile("./clean_laravel-2024-11-07.log")
	if err != nil {
		fmt.Println("Error parsing file", err)
	}
	entries := parseEntries(lines)
	for _, entry := range entries {
		fmt.Printf("Time: %v -> %v \n\n", entry.time, entry.log)
	}*/
	needle := make([]rune, 0)
	needle = append(needle, []rune("Hooligan")...)
	hayStack := make([]rune, 0)
	hayStack = append(hayStack, []rune("Hoola-Hoola girls like Hooligans")...)
	result := Search(needle, hayStack)
	println(result)
	/*result := shiftArray(resultx)
	for i, r := range result {
		fmt.Printf("Index %c: Numder: %d (Unicode %U)\n", i, r, r)
	}*/
}
