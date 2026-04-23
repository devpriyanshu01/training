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

	max := math.MinInt
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

		//call checkPrime fn for answer
	}
	checkPrime(queries, q, max, writer)
}

func checkPrime(queries []int, n, max int, writer io.Writer) {
	//find the factors of each queries
	factors := make([]int, max+1)

	for i := 1; i <= max; i++ {
		for j := i; j <= max; j += i {
			factors[j]++
		}
	}

	for _, val := range queries {
		if factors[val] == 2 {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
