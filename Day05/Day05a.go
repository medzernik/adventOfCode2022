package Day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay05a() {
	f, err := os.Open("./Day05/input_ja.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	//var stacks [9][9]int
	//var stacks [][]rune

	stacks := make([][]rune, 9)

	for scanner.Text() != " 1   2   3   4   5   6   7   8   9 " {
		for i, sym := range scanner.Text() {
			if sym != ' ' && sym != '[' && sym != ']' {
				stacks[(i-1)/4] = append([]rune{sym}, stacks[(i-1)/4]...)
			}
		}
		scanner.Scan()
	}

	scanner.Scan()
	for scanner.Scan() {
		var amount, from, to int

		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &amount, &from, &to)

		//move to box
		for i := 0; i < amount; i++ {

			stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-1])
			fmt.Println(stacks[from-1][len(stacks[from-1])-1])
		}

		fmt.Println(amount, from, to)
	}

	//fmt.Printf("%c", stacks)

}
