package Day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunDay04b() {
	f, err := os.Open("./Day04/input_filip.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ll, lr, rl, rr int
	var count int

	for scanner.Scan() {
		assignment := strings.Split(scanner.Text(), ",")

		for i, j := range strings.Split(assignment[0], "-") {
			if i == 0 {
				ll, _ = strconv.Atoi(j)
			} else {
				lr, _ = strconv.Atoi(j)
			}
		}

		for i, j := range strings.Split(assignment[1], "-") {
			if i == 0 {
				rl, _ = strconv.Atoi(j)
			} else {
				rr, _ = strconv.Atoi(j)
			}
		}

		if (rl <= lr && rr >= ll) || (ll <= rr && lr >= rl) {
			count++
		}

	}
	fmt.Printf("Day 04b result is: \t%d\n", count)
}
