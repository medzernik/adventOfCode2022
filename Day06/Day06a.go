package Day06

import (
	"bufio"
	"log"
	"os"
)

func RunDay06a() {
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
}
