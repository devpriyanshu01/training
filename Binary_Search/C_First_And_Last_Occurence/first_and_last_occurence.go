package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)

		//find first occurence
		first := binarySearchLeftmost(slice, n, x)

		//find last occurence of x
		last := binarySearchLastmost(slice, n, x)

		log.Printf("first - %d, last = %d", first, last)

		if first == n && last == -1 {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintln(writer, first+1, last+1)
		}
	}
}

func binarySearchLastmost(slice []int, n, x int) int {
	l, r := 0, n-1
	m := (l + r) / 2

	var last = -1

	for l <= r {
		if slice[m] == x {
			//check last occurence
			match := m
			if match > last {
				last = match
			}
			l = m + 1
		} else if x > slice[m] {
			l = m + 1
		} else {
			r = m - 1
		}
		m = (l + r) / 2
	}

	return last 
}

// binary search left-most index of target
func binarySearchLeftmost(slice []int, n, x int) int {
	l, r := 0, n-1
	m := (l + r) / 2

	var first = n

	for l <= r {
		if slice[m] == x {
			//check first occurence
			match := m
			if match < first {
				first = match
			}
			r = m - 1
		} else if x > slice[m] {
			l = m + 1
		} else {
			r = m - 1
		}
		m = (l + r) / 2
	}

	return first 
}
