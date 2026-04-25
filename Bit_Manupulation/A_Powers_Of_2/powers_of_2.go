package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	//logic begins here...
	p := 0

	for int(math.Pow(2, float64(p))) <= n {
		fmt.Fprint(writer, int(math.Pow(2, float64(p))), " ")
		p++
	}
	fmt.Fprintln(writer)
}
