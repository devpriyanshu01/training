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

	//testcase
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, x int
		fmt.Fscan(reader, &n)

		slice := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &slice[i])
		}

		fmt.Fscan(reader, &x)

		//core logic begins...
		//what observed is that - if (smallest element of array) <= x <= (largest element of the array), then.
		//there will a case when after performing (n-1) operations, x will be left finally in one case.

		//so let's sort the array
		sort.Ints(slice)
		if x >= slice[0] && x <= slice[n-1] {
			fmt.Fprintln(writer, "Yes")
		}else {
			fmt.Fprintln(writer, "No")
		}
	}
}
