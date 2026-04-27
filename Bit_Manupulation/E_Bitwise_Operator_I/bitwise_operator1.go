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

	var a, b uint
	fmt.Fscan(reader, &a, &b)

	and := a & b
	or := a | b
	xor := a ^ b

	fmt.Fprintln(writer, and, or, xor)
}
