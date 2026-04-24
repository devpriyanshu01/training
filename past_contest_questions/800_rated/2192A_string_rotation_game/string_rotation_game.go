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

	//testcases
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		//string input
		var str string
		fmt.Fscan(reader, &str)

		//core logic begins...
		maxBlock := countBlocks(str, n)
		byteSlice := []byte(str)
		for i := 0; i < n; i++ {
			cyclicRotate(byteSlice, n)
			toString := string(byteSlice)

			count := countBlocks(toString, n)
			if count > maxBlock {
				maxBlock = count
			}
		}

		fmt.Fprintln(writer, maxBlock)
	}
}

func countBlocks(str string, n int) int {
	if n == 0 || n == 1 {
		return n
	}

	count := 1

	for i := 1; i < n; i++ {
		if str[i] != str[i-1] {
			count++
		}
	}
	return count
}

func cyclicRotate(byteSlice []byte, n int) {
	if n < 2 {
		return
	}

	last := byteSlice[0]

	for i := 0; i < n; i++ {
		temp := byteSlice[i]
		byteSlice[i] = last
		last = temp
	}
	byteSlice[0] = last
}
