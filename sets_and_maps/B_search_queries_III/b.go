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

	//n & q
	var n, q int
	fmt.Fscan(reader, &n, &q)

	//slice elements
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	//create a map
	freq := make(map[int]int)
	for i, val := range slice {
		freq[val] = i 
	}

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)

		index, ok := freq[x]
		if ok {
			fmt.Fprintln(writer, index+1)
		}else {
			fmt.Fprintln(writer, -1)
		}
	}
}
