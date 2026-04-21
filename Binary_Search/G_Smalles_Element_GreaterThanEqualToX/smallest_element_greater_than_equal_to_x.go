package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	//input slice of int
	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	//sort the slice for bs
	sort.Ints(slice)

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)

		ans := binarySearchLogic(slice, n, x)
		if ans == math.MaxInt {
			fmt.Fprintln(writer, -1)
		}else {
			fmt.Fprintln(writer, ans)
		}
	}
}

func binarySearchLogic(slice []int, n, x int) int {
	l, r := 0, n-1
	m := (l + r) / 2

	ans := math.MaxInt
	for l <= r {
		if slice[m] >= x {
			ans = slice[m]
			r = m - 1
		} else {
			l = m + 1
		}
		m = (l + r) / 2
	}
	return ans
}
