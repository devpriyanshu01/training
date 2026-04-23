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

	var max = math.MinInt
	queries := make([]int, q)
	j := 0
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)

		if x > max {
			max = x
		}

		queries[j] = x
		j++
	}
	goodNumbers(queries, q, max, writer)
}

func goodNumbers(queries []int, n, max int, writer io.Writer) {
	//find the factors count for each natural numbers till max
	factors := make([]int, max+1)
	for i := 1; i <= max; i++ {
		for j := i; j <= max; j += i {
			factors[j]++
		}
	}

	//logic
	for _, val := range queries {
		factor := factors[val]
		if factors[factor] == 2 {
			fmt.Fprintln(writer, "YES")
		}else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
