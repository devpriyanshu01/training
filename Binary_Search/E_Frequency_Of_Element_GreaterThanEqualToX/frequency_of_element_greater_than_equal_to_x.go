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

	var n, q int
	fmt.Fscan(reader, &n, &q)

	//array elements input
	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	sort.Ints(slice)

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)	

		firstIndex := firstElementGreaterThanEqualToX(slice, n, x)	
		if firstIndex == -1 {
			fmt.Fprintln(writer, 0)
		}else {
			fmt.Fprintln(writer, n-firstIndex)
		}
	}
}

func firstElementGreaterThanEqualToX(slice []int, n, x int) int {
	l, r := 0, n-1
	m := (l+r)/2

	ans := -1

	for l <= r {
		if slice[m] >= x {
			ans = m
			r = m-1
		}else {
			l = m+1
		}
		m = (l+r)/2
	}

	return ans
}
