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

		if n == 2 {
			fmt.Fprintln(writer, n)
		}else if n == 3 {
			fmt.Fprintln(writer, n)
		}else if n & 1 == 0 {
			fmt.Fprintln(writer, 0)
		}else if n & 1 == 1 {
			fmt.Fprintln(writer, 1)
		}
	}
}


