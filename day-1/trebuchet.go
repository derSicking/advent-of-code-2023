package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		first, second := 0, 0
		for i := 0; i < len(line); i++ {
			subline := line[i:]
			firstChar := subline[0]
			if firstChar >= '0' && firstChar <= '9' {
				first = int(firstChar) - '0'
				break
			}
			if strings.HasPrefix(subline, "one") {
				first = 1
				break
			}
			if strings.HasPrefix(subline, "two") {
				first = 2
				break
			}
			if strings.HasPrefix(subline, "three") {
				first = 3
				break
			}
			if strings.HasPrefix(subline, "four") {
				first = 4
				break
			}
			if strings.HasPrefix(subline, "five") {
				first = 5
				break
			}
			if strings.HasPrefix(subline, "six") {
				first = 6
				break
			}
			if strings.HasPrefix(subline, "seven") {
				first = 7
				break
			}
			if strings.HasPrefix(subline, "eight") {
				first = 8
				break
			}
			if strings.HasPrefix(subline, "nine") {
				first = 9
				break
			}
		}
		for i := 0; i < len(line); i++ {
			subline := line[:len(line)-i]
			firstChar := subline[len(subline)-1]
			if firstChar >= '0' && firstChar <= '9' {
				second = int(firstChar) - '0'
				break
			}
			if strings.HasSuffix(subline, "one") {
				second = 1
				break
			}
			if strings.HasSuffix(subline, "two") {
				second = 2
				break
			}
			if strings.HasSuffix(subline, "three") {
				second = 3
				break
			}
			if strings.HasSuffix(subline, "four") {
				second = 4
				break
			}
			if strings.HasSuffix(subline, "five") {
				second = 5
				break
			}
			if strings.HasSuffix(subline, "six") {
				second = 6
				break
			}
			if strings.HasSuffix(subline, "seven") {
				second = 7
				break
			}
			if strings.HasSuffix(subline, "eight") {
				second = 8
				break
			}
			if strings.HasSuffix(subline, "nine") {
				second = 9
				break
			}
		}
		sum += 10*first + second
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println("Sum:", sum)
}
