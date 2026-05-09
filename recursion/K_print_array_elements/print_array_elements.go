package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	//logic starts...
	f(slice, n, writer)
	fmt.Fprintln(writer)

}

func f(slice []int, n int, writer io.Writer) {
	if n == 0 {
		return
	}

	f(slice, n-1, writer)
	fmt.Fprint(writer, slice[n-1], " ")

}
