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

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		exists := false
		var x, y int
		fmt.Fscan(reader, &x, &y)

		for z := x+1; z < y; z++ {
			zmodX := z % x
			ymodZ := y % z

			if zmodX == 0 && ymodZ != 0 {
				fmt.Fprintln(writer, "Yes")
				exists = true
				break
			}
		}

		if !exists {
			fmt.Fprintln(writer, "No")
		}

	}
}
