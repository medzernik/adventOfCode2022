package Day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run_Day01b() {
	fileHandle, _ := os.Open("./Day01/input_ja.txt")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)

	var lineArray []string

	maximum := [3]int{0, 0, 0}
	var total int

	var numValue = 0

	for fileScanner.Scan() {
		lineArray = append(lineArray, fileScanner.Text())
	}

	for _, j := range lineArray {
		tempValue, _ := strconv.Atoi(j)
		numValue += tempValue

		if j == "" {
			for i := range maximum {
				if numValue > maximum[i] {
					maximum[i] = numValue
					break
				}
			}

			numValue = 0

		}
	}

	for _, j := range maximum {
		total += j
	}

	fmt.Printf("Day 1b answer: %d", total)
}
