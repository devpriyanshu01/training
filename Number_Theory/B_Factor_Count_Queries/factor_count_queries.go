package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	queries := make([]int, q)
	max := math.MinInt
	l := 0
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)

		if x > max {
			max = x
		}

		queries[l] = x
		l++
	}

	countFactors(queries, writer, max)
}

func countFactors(queries []int, writer io.Writer, max int) {
	factors := make([]int, max+1)

	for i := 1; i <= max; i++ {
		for j := i; j <= max; j += i {
			factors[j]++
		}
	}

	for _, val := range queries {
		fmt.Fprintln(writer, factors[val])
	}

}
