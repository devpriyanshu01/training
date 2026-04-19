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
		var n, w int
		fmt.Fscan(reader, &n, &w)

		max := w - 1
		ans := 0

		for n > 0 {
			if n == max { //case when a gap is already left & n == max. So taking max
				ans += max - 1 //here make total gap equal to w => lawn mower can leave.
				n -= max
			} else if n > max {
				ans += max
				n -= w
			} else if n < max {
				ans += n
				n -= n
			}
			// n -= 1 //for leaving gap
			// log.Println("n left after gap allocation", n)
			// log.Println("curr answer", ans)
		}
		fmt.Fprintln(writer, ans)
	}
}
