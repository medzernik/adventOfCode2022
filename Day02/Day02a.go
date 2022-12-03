package Day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunDay02a() {
	f, err := os.Open("./Day02/input_filip02.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lineArray []string

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineArray = append(lineArray, scanner.Text())
	}

	var score int

	for i := range lineArray {
		strategy := strings.Fields(lineArray[i])
		opponent := strategy[0]

		switch strategy[1] {
		case "X":
			score += 1

			switch opponent {
			case "C":
				score += 6

			case "A":
				score += 3

			case "B":
				score += 0
			}

		case "Y":
			score += 2

			switch opponent {
			case "A":
				score += 6
			case "B":
				score += 3
			case "C":
				score += 0
			}

		case "Z":
			score += 3

			switch opponent {
			case "B":
				score += 6
			case "C":
				score += 3
			case "A":
				score += 0
			}
		}

	}
	fmt.Printf("Day 02a result is: \t%d\n", score)

}
