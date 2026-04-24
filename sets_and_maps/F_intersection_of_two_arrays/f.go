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

	//n & m
	var n, m int
	fmt.Fscan(reader, &n)

	//taking inputs directly in two maps/sets
	freq1 := make(map[int]struct{})
	freq2 := make(map[int]struct{})

	for i := 0; i < n; i++ {
		var key int
		fmt.Fscan(reader, &key)
		freq1[key] = struct{}{}
	}

	//take m input
	fmt.Fscan(reader, &m)

	for i := 0; i < m; i++ {
		var key int
		fmt.Fscan(reader, &key)
		freq2[key] = struct{}{}
	}

	//core logic begins...

	//slice to store the answer
	var ans []int
	for key := range freq1 {
		_, ok := freq2[key]
		if ok {
			ans = append(ans, key)
		}
	}

	sort.Ints(ans)
	fmt.Fprintln(writer, len(ans))
	for _, val := range ans {
		fmt.Fprint(writer, val, " ")
	}
}
