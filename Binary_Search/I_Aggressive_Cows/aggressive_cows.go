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

	var n, cows int
	fmt.Fscan(reader, &n, &cows)

	stalls := make([]int, n)
	for i := range stalls {
		fmt.Fscan(reader, &stalls[i])
	}

	//maximum possible gap would be - distance of last stall - first stall, i.e. stalls[n-1] - stalls[0]
	l, r := 0, stalls[n-1] - stalls[0]
	m := (l+r)/2

	var ans = -1
	for l <= r {
		if possible(stalls, m, cows) {
			ans = m
			l = m+1
		}else {
			r = m-1
		}
		m = (l+r)/2
	}

	fmt.Fprintln(writer, ans)
}

func possible(stalls []int, gap, cows int) bool {
	lastCowTied := stalls[0]
	cowsTied := 1
	for _, stall := range stalls {
		if stall-lastCowTied >= gap {
			lastCowTied = stall
			cowsTied++
		}
	}
	return (cowsTied >= cows)
}