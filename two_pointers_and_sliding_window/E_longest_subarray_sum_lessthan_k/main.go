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

	var n, k int
	fmt.Fscan(reader, &n, &k)

	//take input - slice elements
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	//core logic
	maxLen := 0
	l := 0
	sum := 0
	for r := 0; r < n; r++ {
		sum += slice[r]
		
		if sum >= k {
			sum -= slice[l]
			l++
		}

		maxLen = max(maxLen, r-l+1)
	}	
	fmt.Fprintln(writer, maxLen)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
