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

	var t int
	fmt.Fscan(reader, &t)

	for x := 0; x < t; x++ {
		var n, k int
		fmt.Fscan(reader, &n, &k)

		a := make([]int, k)
		for z := 0; z < k; z++ {
			fmt.Fscan(reader, &a[z])
		}

		//take b array input
		b := make([]int, n)
		for z := 0; z < n; z++ {
			fmt.Fscan(reader, &b[z])
		}

		//preparations for the problem
		aMap := make(map[int]int)
		for ind, val := range a {
			aMap[ind+1] = val
		}

		bMap := make(map[int]int)
		for _, val := range b {
			bMap[val]++
		}

		//array to store ans
		var ans []int

		//logic begins...
		i := 0
		for i < 1000 {
			for j := 0; j < n; j++ {
				currPriority := b[i]

				if currPriority == (k + 1) {
					continue
				}

				maxAllowed := aMap[currPriority+1]
				currCount := bMap[currPriority+1]

				if currPriority == k {
					b[j]++
					i++
					ans = append(ans, j+1)
				} else if currCount < maxAllowed {
					b[j]++
					i++
					ans = append(ans, j+1)
					bMap[currPriority+1]++
					bMap[currPriority]--
				}

				if i >= 1000 {
					break
				}

				if j == n-1 {
					j = 0
				}

				fmt.Println("curr array -> ", b)

			}
			break
		}
		if i >= 1000 {
			fmt.Fprintln(writer, -1)
		}
		fmt.Fprintln(writer, len(ans))
		for _, val := range ans {
			fmt.Fprint(writer, val, " ")
		}
		fmt.Fprintln(writer)
	}
}
