package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Number int

const (
	zero Number = iota
	one
	two
	three
	four
	five
	six
	seven
	eight
	nine
)

func (n Number) String() string {
	return [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}[n]
}

func sumNumbersFromLine(rd bufio.Reader) int {
	fileSum := 0

	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
		}

		occurences := make(map[int]int)

		for i := zero; i <= nine; i++ {
			numWord := Number.String(i)
			re := regexp.MustCompile(numWord)
			indices := re.FindAllStringIndex(line, -1)
			for _, index := range indices {
				startIndex := index[0]
				occurences[startIndex] = int(i)
			}
		}

		re := regexp.MustCompile("[0-9]")
		indices := re.FindAllStringIndex(line, -1)
		for _, index := range indices {
			startIndex := index[0]
			expression := line[startIndex : startIndex+1]
			number, err := strconv.Atoi(expression)
			if err == nil {
				occurences[startIndex] = number
			}
		}

		keys := make([]int, 0, len(occurences))
		for k := range occurences {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		calibValue := strconv.Itoa(occurences[keys[0]]) + strconv.Itoa(occurences[keys[len(keys)-1]])
		calibrationValue, err := strconv.Atoi(calibValue)
		if err != nil {
			fmt.Println(err)
		}
		fileSum += calibrationValue
	}
	return fileSum
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	rd := bufio.NewReader(file)

	fmt.Println(sumNumbersFromLine(*rd))
}
