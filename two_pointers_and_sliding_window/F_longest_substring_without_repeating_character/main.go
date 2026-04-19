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

	var str string
	fmt.Fscan(reader, &str)

	//main logic
	l := 0
	maxLen := 0

	sub := make(map[byte]int)
	for r := 0; r < n; r++ {
		sub[str[r]]++
		val, _ := sub[str[r]]
		for val > 1 {
			sub[str[l]]--
			v, _ := sub[str[l]]
			if v == 0 {
				delete(sub, str[l])
			}
			val, _ = sub[str[r]]
			l++
		}
		maxLen = max(maxLen, len(sub)) 
	}
	fmt.Fprintln(writer, maxLen)
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
