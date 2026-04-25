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

	set := make(map[int]struct{})
	for i := 2; i*i <= n; i++ {
		if n % i == 0 {
			for n % i == 0 {
				set[i] = struct{}{}
				n = n/i
			}
		}
	}

	if n > 1 {
		set[n] = struct{}{}
	}

	fmt.Fprintln(writer, len(set))
}
