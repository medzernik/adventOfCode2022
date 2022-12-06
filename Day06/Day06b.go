package Day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay06b() {
	f, err := os.Open("./Day06/input_filip.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()

	for i := range scanner.Text() {
		var characters []int

		for j := 0; j < 14; j++ {
			if contains(characters, int(scanner.Text()[i+j])) == false {
				characters = append(characters, int(scanner.Text()[i+j])) //if characters array doesn't already contain a character, add it
			}
		}

		if len(characters) == 14 {
			fmt.Println(i + 14)
			break
		}
	}
}

//the function is already in Day06a, just putting it here as well
//func contains(s []int, i int) bool {
//	for _, v := range s {
//		if v == i {
//			return true
//		}
//	}
//
//	return false
//}
