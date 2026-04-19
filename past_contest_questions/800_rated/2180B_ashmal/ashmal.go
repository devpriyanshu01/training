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

		slice := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &slice[i])
		}

		ans := slice[0]

		for i := 1; i < n; i++ {
			if ans <= slice[i] {
				ans = ans + slice[i]
			}else {
				ans = slice[i] + ans
			}
		}
		//save to writer

		fmt.Fprintln(writer, ans)
	}	
}
