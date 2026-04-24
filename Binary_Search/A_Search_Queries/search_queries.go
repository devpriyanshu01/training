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

	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)
		
		ans := binarySearch(slice, 0, n-1, x)
		fmt.Fprintln(writer, ans)
	
	}
}

func binarySearch(arr []int, l, r, x int) int {
	m := (l+r)/2

	for l <= r {
		if arr[m] == x {
			return m+1
		}else if arr[m] > x {
			r = m-1
		}else {
			l = m+1
		}
		m = (l+r)/2
	}
	return -1
}
