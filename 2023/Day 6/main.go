package main

import "fmt"

type Race struct {
	time     int
	distance int
}

func main() {
	inputPartOne := map[int]Race{
		1: {45, 305},
		2: {97, 1062},
		3: {72, 1110},
		4: {95, 1695},
	}

	inputPartTwo := map[int]Race{
		1: {45977295, 305106211101695},
	}

	totalWaysToWin := 1

	for k := range inputPartOne {
		wayToWin := 0
		time := inputPartOne[k].time
		distance := inputPartOne[k].distance
		for i := 1; i < time; i++ {
			speed := i
			if speed*(time-speed) > distance {
				wayToWin++
			}
		}
		totalWaysToWin *= wayToWin
	}

	for k := range inputPartTwo {
		wayToWinPartTwo := 0
		time := inputPartTwo[k].time
		distance := inputPartTwo[k].distance
		for i := 1; i < time; i++ {
			speed := i
			if speed*(time-speed) > distance {
				wayToWinPartTwo++
			}
		}
		fmt.Println(wayToWinPartTwo)
	}
	fmt.Println(totalWaysToWin)
	
	
}