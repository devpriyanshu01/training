package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//take size input
	n := 0
	fmt.Fscan(reader, &n)

	//fill array element
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &array[i])
	}

	//create prefix array
	prefix := make([]int, n)
	prefix[0] = array[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + array[i]
	}

	//logic
	sum := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == 0 {
				sum += prefix[j]
			} else {
				sum += prefix[j] - prefix[i-1]
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(writer, sum)
	writer.Flush()
}
