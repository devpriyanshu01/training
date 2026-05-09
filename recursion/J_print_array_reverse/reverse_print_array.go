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

    array := make([]int, n)
    for i := range n {
    	fmt.Fscan(reader, &array[i])
    }

    f(array, n, writer)
    fmt.Fprintln(writer)
}

func f(array []int, n int, writer io.Writer) {
	if n == 0 {
		return
	}

	fmt.Fprint(writer, array[n-1], " ")
	f(array, n-1, writer)

}
