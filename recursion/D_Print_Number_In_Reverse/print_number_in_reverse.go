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

	if n == 0 {
		fmt.Fprint(writer, n)
	}else {

		f(n, writer)
	}


	fmt.Fprintln(writer)

}

func f(n int, writer io.Writer) {
	if n == 0 {
		return
	}

	fmt.Fprint(writer, n%10)
	f(n/10, writer)
}
