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

	var n, k int
	fmt.Fscan(reader, &n, &k)

	
	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}
	
	if k > n {
		fmt.Fprintln(writer, -1)
		return
	}

	//minimum pages that can be assigned will the max element in the slice
	max := math.MinInt
	totalPages := 0
	for _, val := range slice {
		if val > max {
			max = val
		}
		totalPages += val
	}

	l, r := max, totalPages
	currPages := (l+r)/2

	var ans = -1
	for l <= r {
		if possible(slice, n, currPages, k) {
			r = currPages-1
			ans = currPages
		}else {
			l = currPages+1
		}
		currPages = (l+r)/2
	}

	fmt.Fprintln(writer, ans)
}

func possible(slice []int, n, pages, students int) bool {
	studentsAssigned := 0

	totalPages := 0
	for _, val := range slice {
		totalPages += val	
		if totalPages > pages {
			studentsAssigned++
			totalPages = val
		}
	}
	studentsAssigned++

	return studentsAssigned <= students
}
