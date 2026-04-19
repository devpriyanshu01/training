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

	var n, q int32
	fmt.Fscan(reader, &n, &q)

	slice := make([]int32, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	for i := 0; i < int(q); i++ {
		var x int32
		fmt.Fscan(reader, &x)

		ind := BinarySearchDecSorted(slice, n, 0, n-1, x)

		fmt.Fprintln(writer, ind)
	}

}

func BinarySearchDecSorted(slice []int32, n, l, r, x int32) int32 {
	m := (l+r)/2

	for l <= r {
		if slice[m] == x {
			return m+1
		}else if x > slice[m] {
			r = m-1
		}else {
			l = m+1
		}
		m = (l+r)/2
	}

	return -1
}
