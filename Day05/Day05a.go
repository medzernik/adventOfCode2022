package Day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay05a() {
	f, err := os.Open("./Day05/input_filipino.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	stacks := make([][]rune, 9)

	for scanner.Text() != " 1   2   3   4   5   6   7   8   9" { //there should be an extra space after 9, but you can try what happens if you add it yourself :))))
		for i, sym := range scanner.Text() {
			if sym != ' ' && sym != '[' && sym != ']' {
				stacks[(i-1)/4] = append([]rune{sym}, stacks[(i-1)/4]...)
			}
		}
		fmt.Println(stacks)
		scanner.Scan()
	}

	scanner.Scan()
	for scanner.Scan() {
		var amount, from, to int

		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &amount, &from, &to)

		//move to box
		for i := 0; i < amount; i++ {
			//there should be proper stack implementation methods here
			ti, fi, l := to-1, from-1, len(stacks[from-1])   //"ti" and "fi" stand for "to index" and "from index"
			stacks[ti] = append(stacks[ti], stacks[fi][l-1]) //adds last item from "from" stack to "to" stack
			stacks[fi] = stacks[fi][:l-1]                    //removes last item from "from" stack
			//fmt.Println(stacks[fi])
		}

		fmt.Println(amount, from, to)

	}

	for _, box := range stacks {
		fmt.Print(string(box[len(box)-1])) //prints last of each stack
	}

}
