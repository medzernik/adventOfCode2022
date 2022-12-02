package Day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunDay02b() {
	f, err := os.Open("./Day02/input_filip02.txt")
	if err != nil {
		log.Fatal(err)
	}


	var lineArrays []string

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineArrays = append(lineArrays, scanner.Text())
	}

	var score int

	for i := range lineArrays {
		strategy := strings.Fields(lineArrays[i])
		opponent := strategy[0]

		switch strategy[1] {
		case "X":
			score += 0

			switch opponent {
			case "A":
				score += 3
			case "B":
				score += 1
			case "C":
				score += 2
			}

		case "Y":
			score += 3

			switch opponent {
			case "A":
				score += 1
			case "B":
				score += 2
			case "C":
				score += 3
			}

		case "Z":
			score += 6

			switch opponent {
			case "A":
				score += 2
			case "B":
				score += 3
			case "C":
				score += 1
			}
		}

	}
	fmt.Printf("Day 02b result is: %d\n", score)

}
