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
		fmt.Println(j)
		fmt.Println(leftPart)
		fmt.Println(rightPart)

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

		//splitPart := strings.Split(lineArray[i], "")
		//
		//rightPart := splitPart[len(splitPart)/2:]
		//leftPart := splitPart[:len(splitPart)/2]
		//fmt.Print(leftPart)
		//fmt.Print(" | ")
		//fmt.Print(rightPart)
		//fmt.Println()

		//fmt.Println(test)
	}

	fmt.Println(prioritiesSum)

	fmt.Println("INTEGERS FOLLOW")
	fmt.Println(int('a') - 96)
	fmt.Println(int('z') - 96)
	fmt.Println(int('A') - 38)
	fmt.Println(int('Z') - 38)
}
