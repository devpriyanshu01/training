package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//size of array
	n := 0
	fmt.Fscan(reader, &n)

	//no. of queries
	q := 0
	fmt.Fscan(reader, &q)

	//fill array with elements
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &array[i])
	}

	//prefix array
	prefix := make([]int, n)
	square := math.Pow(float64(array[0]), 2)
	prefix[0] = int(square)
	for i := 1; i < n; i++ {
		squareSum := prefix[i-1] + int(math.Pow(float64(array[i]), 2))
		prefix[i] = squareSum
	}

	//create a writer
	writer := bufio.NewWriter(os.Stdout)
	ans := 0
	for i := 0; i < q; i++ {
		l, r := 0, 0
		fmt.Fscan(reader, &l, &r)
		// fmt.Fscan(reader, &r)

		if l == 1 {
			ans = prefix[r-1]
		} else {
			ans = prefix[r-1] - prefix[l-2]
		}
		fmt.Fprintln(writer, ans)
	}

	writer.Flush()

}
