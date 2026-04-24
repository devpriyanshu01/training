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

	//number of testcases
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var x int
		fmt.Fscan(reader, &x)

		if x % 9 == 0 {
			fmt.Fprintln(writer, 10)
		}else {
			fmt.Fprintln(writer, 0)
		}
	}
}
