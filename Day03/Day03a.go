package Day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func RunDay03a() {
	f, err := os.Open("./Day03/input_ja.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lineArray []string

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineArray = append(lineArray, scanner.Text())
	}

	//fmt.Println(int('A'))
	var prioritiesSum int

	for _, j := range lineArray {

		leftPart := j[:len(j)/2]
		rightPart := j[len(j)/2:]
		//fmt.Println(j)
		//fmt.Println(leftPart)
		//fmt.Println(rightPart)

		for i := range rightPart {
			if strings.Contains(string(rightPart), string(leftPart[i])) {
				item := rune(leftPart[i])

				if unicode.IsLower(item) {
					prioritiesSum += int(item) - 96
				} else if unicode.IsUpper(item) {
					prioritiesSum += int(item) - 38
				}
				break
			}
		}
	}

	fmt.Printf("Day 03a result is: \t%d\n", prioritiesSum)
}
