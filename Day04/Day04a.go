package Day04

import (
	"bufio"
	"log"
	"os"
)

func RunDay04a() {
	f, err := os.Open("./Day04/input_ja.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lineArray []string

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineArray = append(lineArray, scanner.Text())
	}
}
