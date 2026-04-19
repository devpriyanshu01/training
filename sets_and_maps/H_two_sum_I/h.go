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

	//n & k
	var n, k int
	fmt.Fscan(reader, &n, &k)

	//slice input
	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	//core logic begins
	set := make(map[int]int)
	for _, val := range slice {
		set[val]++
	}

	found := false
	for _, val := range slice {
		reqd := k - val

		count, exist := set[reqd]
		if exist {
			if val == reqd {
				if count > 1 {
					fmt.Fprintln(writer, "TRUE")
					found = true
					break
				}
			} else if val != reqd {
				fmt.Fprintln(writer, "TRUE")
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Fprintln(writer, "FALSE")
	}
}
