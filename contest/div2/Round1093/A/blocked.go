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
	
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		slice := make([]int, n)
		for i := range slice {
			fmt.Fscan(reader, &slice[i])
		}

		sort.Ints(slice)

		//check if it contains duplicates
		mymap := make(map[int]int)

		for _, val := range slice {
			mymap[val]++
		}

		done := false
		for _, val := range mymap {
			if val > 1 {
				fmt.Fprintln(writer, -1)
				done = true
			}
		}

		if !done {
			for i := n-1; i >= 0; i-- {
				fmt.Fprint(writer, slice[i], " ")
			}	
			fmt.Fprintln(writer)
		}

	}
}
