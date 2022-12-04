package Day04

import (
	"bufio"
	"fmt"
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

		if (rl >= ll && rr <= lr) || (ll >= rl && lr <= rr) {
			count++
		}

	}
}
