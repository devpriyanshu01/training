package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	countFactors(n, writer)

}

func countFactors(n int, writer io.Writer) {
	factors := make([]int, n+1)

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			factors[j]++
		}
	}

	primes := 0
	for i, val := range factors {
		if val == 2 {
			fmt.Fprint(writer, i, " ")
			primes++
		}
	}

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, primes)
}
