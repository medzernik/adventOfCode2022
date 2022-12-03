package Day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunDay01a() {
	fileHandle, _ := os.Open("./Day01/input_ja.txt")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)

	var lineArray []string

	var maximum int
	var num_value = 0

	for fileScanner.Scan() {
		lineArray = append(lineArray, fileScanner.Text())
	}

	for _, j := range lineArray {
		temp_value, _ := strconv.Atoi(j)
		num_value += temp_value

		if j == "" {
			if num_value > maximum {
				maximum = num_value
			}
			num_value = 0

		}
	}

	fmt.Printf("Day 01a result is: \t%d\n", maximum)
}
