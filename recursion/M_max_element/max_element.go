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

	var t int
	fmt.Fscan(reader, &t)

	slice := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	ans := f(slice, t)
	fmt.Fprintln(writer, ans)
}

func f(slice []int, n int) int {
	if n == 1 {
		return slice[n-1]
	}

	small := f(slice, n-1)
	ans := max(small, slice[n-1])
	return ans
}