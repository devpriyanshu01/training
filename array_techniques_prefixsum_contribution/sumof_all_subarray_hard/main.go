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

	n := 0
	fmt.Fscan(reader, &n)

	//array elements
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	//logic to find sum of all subarrays
	sum := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum += arr[j]
		}
	}

	fmt.Fprintln(writer, sum)
}
