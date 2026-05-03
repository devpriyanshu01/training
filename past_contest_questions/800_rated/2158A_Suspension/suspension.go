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

		var y, r int
		fmt.Fscan(reader, &y, &r)

		//logic begins here...
		var ans int
		//work on red cards
		ans += r
		
		//work on yellow card
		ans += y/2

		if ans >= n {
			fmt.Fprintln(writer, n)
		}else {
			fmt.Fprintln(writer, ans)
		}
		
		
	}
}
