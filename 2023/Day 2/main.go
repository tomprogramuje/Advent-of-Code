package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func sumOfIDs(rd bufio.Reader) (int, int) {
	idSum := 0
	powerSum := 0

	for {
		line, _, err := rd.ReadLine()
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
			if pattern.MatchString(string(line)) {
				notAllowed = false
				break
			}
		}
		if notAllowed {
			gameID := regexp.MustCompile(`Game (\d+):`)
			match := gameID.FindStringSubmatch(string(line))
			value, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println(err)
			}
			idSum += value
		}
		red := regexp.MustCompile(`\b(\d+) red\b`)
		allRed := red.FindAllStringSubmatch(string(line), -1)
		green := regexp.MustCompile(`\b(\d+) green\b`)
		allGreen := green.FindAllStringSubmatch(string(line), -1)
		blue := regexp.MustCompile(`\b(\d+) blue\b`)
		allBlue := blue.FindAllStringSubmatch(string(line), -1)

		var allRedCubes, allGreenCubes, allBlueCubes []int
		powerGame := 1	

		for i := range allRed {
			num := strings.Split(allRed[i][0], " ")
			numParsed, err := strconv.Atoi(num[0])
			if err != nil {
				fmt.Println(err)
			}
			allRedCubes = append(allRedCubes, numParsed)
		}
		sort.Ints(allRedCubes)
		powerGame *= allRedCubes[len(allRedCubes)-1]

		for i := range allGreen {
			num := strings.Split(allGreen[i][0], " ")
			numParsed, err := strconv.Atoi(num[0])
			if err != nil {
				fmt.Println(err)
			}
			allGreenCubes = append(allGreenCubes, numParsed)
		}
		sort.Ints(allGreenCubes)
		powerGame *= allGreenCubes[len(allGreenCubes)-1]

		for i := range allBlue {
			num := strings.Split(allBlue[i][0], " ")
			numParsed, err := strconv.Atoi(num[0])
			if err != nil {
				fmt.Println(err)
			}
			allBlueCubes = append(allBlueCubes, numParsed)
		}
		sort.Ints(allBlueCubes)
		powerGame *= allBlueCubes[len(allBlueCubes)-1]

		powerSum += powerGame
	}
	return idSum, powerSum
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
