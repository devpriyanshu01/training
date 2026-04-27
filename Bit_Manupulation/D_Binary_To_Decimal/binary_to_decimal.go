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

	var str string
	fmt.Fscan(reader, &str)

	n := len(str)

	var num int
	for i := n - 1; i >= 0; i-- {
		if str[i] == '1' {
			num += int(math.Pow(2, float64(n-i-1)))
		}
	}
	fmt.Fprintln(writer, num)
}
