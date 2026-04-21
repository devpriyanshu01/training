package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//input n & q
	var n, q int
	fmt.Fscan(reader, &n, &q)

	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	//sort the slice so that binary search can be applied
	sort.Ints(slice)

	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)

		if x > slice[n-1] || x > y {
			fmt.Fprintln(writer, 0)
		} else {
			indexX := binarySearch(slice, n, x)
			indexY := binarySearch2(slice, n, y)

			ans := indexY - indexX + 1
			fmt.Fprintln(writer, ans)
		}
	}
}

func binarySearch(slice []int, n, x int) int {
	l, r := 0, n-1
	m := (l + r) / 2

	index := -1

	for l <= r {
		if slice[m] >= x {
			index = m
			r = m - 1
		} else {
			l = m + 1
		}
		m = (l + r) / 2
	}
	return index
}

func binarySearch2(slice []int, n, y int) int {
	l, r := 0, n-1
	m := (l + r) / 2

	index := -1
	for l <= r {
		if slice[m] <= y {
			index = m
			l = m + 1
		} else {
			r = m - 1
		}
		m = (l + r) / 2
	}
	return index
}
