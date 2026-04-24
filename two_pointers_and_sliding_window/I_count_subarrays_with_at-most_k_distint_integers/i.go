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

	//n & k input
	var n, k int
	fmt.Fscan(reader, &n, &k)

	//integer slice input
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	//core logic begins...
	i, ans := 0, 0

	freq := make(map[int]int)
	for j := 0; j < n; j++ {
		freq[slice[j]]++

		for len(freq) > k {
			freq[slice[i]]--
			val := freq[slice[i]]
			if val == 0 {
				delete(freq, slice[i])
			}
			i++
		}

		ans += j - i + 1
	}

	fmt.Fprintln(writer, ans)
}
