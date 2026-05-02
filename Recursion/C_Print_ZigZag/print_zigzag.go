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

	f(n, writer)

}

func f(n int, writer io.Writer) {
	if n == 1 {
		fmt.Fprintln(writer, n)
		return
	}

	fmt.Fprintln(writer, n)
	f(n-1, writer)
	fmt.Fprintln(writer, n)
	
}
