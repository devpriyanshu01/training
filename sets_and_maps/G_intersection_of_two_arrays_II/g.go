package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//n & m input
	var n, m int

	//n input
	fmt.Fscan(reader, &n)

	//first array input
	slice1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice1[i])
	}

	//m input
	fmt.Fscan(reader, &m)

	//input second array into a map
	freq := make(map[int]int)
	for i := 0; i < m; i++ {
		var key int
		fmt.Fscan(reader, &key)
		freq[key]++
	}

	//core logic begins
	//creating a slice to store the answer
	var ans []int
	for i := 0; i < n; i++ {
		val := slice1[i]

		count, ok := freq[val]
		if ok {
			ans = append(ans, val)
			freq[val]--
			count = freq[val] //get the updated count now
			if count == 0 {
				delete(freq, val)
			}
		}
	}

	//sort the ans as required in the question
	sort.Ints(ans)
	fmt.Fprintln(writer, len(ans))

	for _, val := range ans {
		fmt.Fprint(writer, val, " ")
	}
}
