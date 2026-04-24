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

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		slice := make([]int, n)
		for i := range slice {
			fmt.Fscan(reader, &slice[i])
		}

		//sort the slice
		sort.Ints(slice)
		
		//maximum possible average will the value of largest element in the slice.
		fmt.Fprintln(writer, slice[n-1])
	}
}
