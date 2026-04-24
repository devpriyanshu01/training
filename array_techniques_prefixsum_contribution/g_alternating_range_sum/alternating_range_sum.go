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

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// create prefix array
	prefix := make([]int, n)

	for i := 0; i < n; i++ {
		val := arr[i]
		if i%2 == 1 {
			val = -val
		}

		if i == 0 {
			prefix[i] = val
		} else {
			prefix[i] = prefix[i-1] + val
		}
	

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)

		l-- // convert to 0-based
		r--

		res := prefix[r]
		if l > 0 {
			res -= prefix[l-1]
		}

		// adjust sign if l is odd
		if l%2 == 1 {
			res = -res
		}

		fmt.Fprintln(writer, res)
	}
}
