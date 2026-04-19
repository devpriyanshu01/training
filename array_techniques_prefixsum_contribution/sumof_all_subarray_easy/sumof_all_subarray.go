package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//size of array
	n := 0
	fmt.Fscan(reader, &n)

	//input array
	arr := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &arr[i])
	}

	//logic begins here
	sum := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(writer, sum)
	writer.Flush()
}
