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

	//n, k
	var n, k int
	fmt.Fscan(reader, &n, &k)

	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	//if you encounter any problem that says - find minimum of maximum or maximum of minimu. Be sure
	//that this is a binary search problem.
	
	//minimum time required to complete the painting when painters are unlimited will be max element in slice.
	//And the maximum time to complete the painting would be when there is only painter. So in this case, 
	//time required to paint will sum of all the elements in the slice.

	//find maximum in the slice
	max := math.MinInt
	sum := 0
	for _, val := range slice {
		if val > max {
			max = val
		}
		sum += val
	}

	//let's start applying binary search
	l := max
	r := sum	
	m := (l+r)/2

	ans := -1
	for l <= r {
		if possible(slice, n, m, k) {
			ans = m
			r = m-1
		}else {
			l = m+1
		}
		m = (l+r)/2
	}

	fmt.Fprintln(writer, ans)
}

func possible(slice []int, n, givenTime, availablePainters int) bool {
	requiredPainters := 0

	sum := 0
	for _, val := range slice {
		sum += val
		if sum > givenTime {
			requiredPainters++
			sum = val
		} 
	}
	requiredPainters++
	
	return requiredPainters <= availablePainters
}
