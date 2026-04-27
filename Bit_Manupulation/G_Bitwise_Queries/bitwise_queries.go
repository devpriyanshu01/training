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

	var n, q int
	fmt.Fscan(reader, &n, &q)

	for i := 0; i < q; i++ {
		var t, i int
		fmt.Fscan(reader, &t, &i)

		mask := 1 << i

		switch t {
		case 1:
			//check if ith bit is 1 or 0
			if n & mask == 0 {
				fmt.Fprintln(writer, "NO")
			}else {
				fmt.Fprintln(writer, "YES")
			}
		case 2:
			//set the ith bit
			n = n | mask
			fmt.Fprintln(writer, n)
		case 3:
			//unset the ith bit
			imask := mask
			imask = ^imask
			n = n & imask
			fmt.Fprintln(writer, n)
		case 4:
			//toggle the ith bit
			n = n ^ mask
			fmt.Fprintln(writer, n)
		}
	}
}
