package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for z := 0; z < t; z++ {
		var n int
		fmt.Fscan(reader, &n)

		nums := make([]int, n)
		for i := range nums {
			fmt.Fscan(reader, &nums[i])
		}

		//logic starts here...
		//Alogrithm:
		//1. find the index of the maximum element & swap this posn with 0.
		maxVal, maxInd := math.MinInt, -1	
		for i, num := range nums {
			if num > maxVal {
				maxVal = num
				maxInd = i
			}
		}

		//swap the max element & first element
		temp := nums[0]
		nums[0] = nums[maxInd]
		nums[maxInd] = temp

		for _, num := range nums {
			fmt.Fprint(writer, num, " ")
		}
		fmt.Fprintln(writer)
	}

}
