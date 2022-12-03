package Day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func RunDay03b() {
	f, err := os.Open("./Day03/input_filip03.txt")
	if err != nil {
		log.Fatal(err)
	}

	//var lineArray []string

	defer f.Close()

	scanner := bufio.NewScanner(f)

	//for scanner.Scan() {
	//	lineArray = append(lineArray, scanner.Text())
	//}

	//fmt.Println(int('A'))
	var prioritiesSum int

	for scanner.Scan() {

		firstSet := scanner.Text()
		scanner.Scan()
		secondSet := scanner.Text()
		scanner.Scan()
		thirdSet := scanner.Text()

		fmt.Println(firstSet)
		fmt.Println(secondSet)
		fmt.Println(thirdSet)

		for i := range firstSet {
			if strings.Contains(secondSet, string(firstSet[i])) && strings.Contains(thirdSet, string(firstSet[i])) {
				item := rune(firstSet[i])
				fmt.Println(string(item))

				if unicode.IsLower(item) {
					prioritiesSum += int(item) - 96
				} else if unicode.IsUpper(item) {
					prioritiesSum += int(item) - 38
				}
				break
			}
		}

	}

	fmt.Printf("Day 03b result is: \t%d\n", prioritiesSum)
}
