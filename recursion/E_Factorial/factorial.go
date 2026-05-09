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

	ans := f(n)
	fmt.Fprintln(writer, ans)
}

func f(n int) int {
	if n == 0 {
		return 1
	}

	return n * f(n-1)
}
