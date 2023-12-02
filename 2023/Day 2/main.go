package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func sumOfIDs(rd bufio.Reader) int {
	idSum := 0

	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
		}

		patterns := []*regexp.Regexp{
			regexp.MustCompile(`\b(1[3-9]|[2-9][0-9])\sred\b`),
			regexp.MustCompile(`\b(1[4-9]|[2-9][0-9])\sgreen\b`),
			regexp.MustCompile(`\b(1[5-9]|[2-9][0-9])\sblue\b`),
		}
		notAllowed := true
		for _, pattern := range patterns {
			if pattern.MatchString(line) {
				notAllowed = false
				break
			}
		}
		if notAllowed {
			fmt.Println("valid game")
			gameID := regexp.MustCompile(`Game (\d+):`)
			match := gameID.FindStringSubmatch(line)
			fmt.Println(match[1])
			value, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println(err)
			}
			idSum += value
		}
	}
	return idSum
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	rd := bufio.NewReader(file)
	fmt.Println(sumOfIDs(*rd))
}
