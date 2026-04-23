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
	factors := make([]int, n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j += i {
			factors[j]++
		}
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)

		primes := 0
		for j := l; j <= r; j++ {
			if factors[j] == 2 {
				primes++
			}
		}

		fmt.Fprintln(writer, primes)

	}
}
