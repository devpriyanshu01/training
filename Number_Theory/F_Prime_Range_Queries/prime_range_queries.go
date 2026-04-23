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

	//generate the factors count
	factors := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			factors[j]++
		}
	}

	prefix := make([]int, n+1)
	count := 0

	for i, val := range factors {
		if i == 0 || i == 1 {
			prefix[i] = 0
		} else {
			if val == 2 {
				count++
				prefix[i] = count
			} else {
				prefix[i] = count
			}
		}
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)

		ans := prefix[r] - prefix[l-1]
		fmt.Fprintln(writer, ans)
	}
}
