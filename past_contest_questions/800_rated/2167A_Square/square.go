package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		//input sides
		slice := make([]int, 4)
		for j := range slice {
			fmt.Fscan(reader, &slice[j])
		}

		equalSides := true
		side := slice[0]
		for _, val := range slice {
			if val != side {
				fmt.Fprintln(writer, "No")
				equalSides = false
				break
			}
		}

		if equalSides {
			fmt.Fprintln(writer, "Yes")
		}
	}

}
