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

	fmt.Fprintln(writer, f(n))

}

func f(n int) int {	
	if n == 1 || n == 2 {
		return n-1
	}
	return f(n-1) + f(n-2)
}
