package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func sumNumbersFromLine(rd bufio.Reader) int {
	fileSum := 0
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			return fileSum
		}
		if err != nil {
			fmt.Println("Error:", err)
		}
	
		re := regexp.MustCompile("[0-9]")
		nums := re.FindAllString(line, -1)	
		firstNum := nums[0]
		secondNum := nums[len(nums)-1]
		
		calibValue := firstNum + secondNum
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
