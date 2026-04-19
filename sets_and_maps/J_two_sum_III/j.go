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

	//prepare 2 maps to help us to search quickly
	freq := make(map[int]int)

	//create freq map for handling duplicate elements
	for _, val := range slice {
		freq[val]++
	}

	twoSumPairs := 0
	//main logic
	for _, val := range slice {
		if freq[val] == 0 {
			continue
		}
		reqd := k - val

		if freq[reqd] > 0 && freq[val] > 0 {
			if reqd == val {
				if freq[reqd] < 2 {
					continue
				}
			}

			freq[reqd]--
			freq[val]--
			twoSumPairs++
		}
	}
	fmt.Fprintln(writer, twoSumPairs)
}
