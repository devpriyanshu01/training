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

	var n int
	fmt.Fscan(reader, &n)

	//slice input
	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	//logic begins here...
	//and operation
	and := slice[0]
	for i := 1; i < n; i++ {
		and &= slice[i]
	}

	//or operation
	or := slice[0]
	for i := 1; i < n; i++ {
		or |= slice[i]
	}

	//xor operation
	xor := slice[0]
	for i := 1; i < n; i++ {
		xor ^= slice[i]
	}

	fmt.Fprintln(writer, and, or, xor)
}
