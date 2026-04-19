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
		var n int
		fmt.Fscan(reader, &n)

		if n&1 == 1 {
			fmt.Fprintln(writer, 0)
		}else {
			//when given n is even
			chicken := n / 2
			var ways int
			for chicken >= 0 {
				ways++
				chicken -= 2
			}
			fmt.Fprintln(writer, ways)
		}

	}
}
