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

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var s string
	fmt.Fscan(reader, &s)

	//core logic of the code begins here...
	l := 0
	freq := make(map[byte]int)
	maxLen := 0

	for r := 0; r < n; r++ {
		freq[s[r]]++

		for len(freq) > k {
			freq[s[l]]--
			v := freq[s[l]]
			if v == 0 {
				delete(freq, s[l])
			}
			l++
		}
		maxLen = max(maxLen, r-l+1)
	}
	fmt.Fprintln(writer, maxLen)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
