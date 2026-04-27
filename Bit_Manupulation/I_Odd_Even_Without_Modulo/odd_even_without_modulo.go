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

	var n int64
	fmt.Fscan(reader, &n)

	if n&1 == 1 {
		fmt.Fprintln(writer, "Odd")
	} else {
		fmt.Fprintln(writer, "Even")
	}
}
