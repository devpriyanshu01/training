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

	var n, x int
	fmt.Fscan(reader, &n, &x)

	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	//logic begins here
	//whenever we see a subarray kind of problem we should think of using prefix sum.

	prefix := make([]int, n)
	for i, val := range slice {
		if i == 0 {
			prefix[i] = val
		} else {
			prefix[i] = prefix[i-1] + val
		}
	}

	//at any point in prefix array if prefix[i] - target exists meaning has a subarray with sum k
	set := make(map[int]int)
	for i, val := range prefix {
		if val == x {
			fmt.Fprintln(writer, 1, i+1)
			return
		}

		ind, exists := set[val-x]
		if exists {
			fmt.Fprintln(writer, ind+2, i+1)
			return
		} else {
			set[val] = i
		}
	}

	fmt.Fprintln(writer, -1)

}
