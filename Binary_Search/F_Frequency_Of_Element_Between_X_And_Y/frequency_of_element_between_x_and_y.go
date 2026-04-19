package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//input n & q
	var n, q int
	fmt.Fscan(reader, &n, &q)

	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	//sort the slice so that binary search can be applied
	sort.Ints(slice)

	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)

		log.Printf("for x = %d & y = %d", x, y)
		if x > slice[n-1] {
			fmt.Fprintln(writer, 0)
		} else {
			indexX := binarySearch(slice, n, x)
			indexY := binarySearch2(slice, n, y)

			log.Printf("indexX = %d & indexY = %d", indexX, indexY)

			ans := indexY - indexX + 1
			fmt.Fprintln(writer, ans)
		}
	}
}

func binarySearch(slice []int, n, x int) int {
	l, r := 0, n-1
	m := (l + r) / 2

	index := -1

	for l <= r {
		fmt.Println("m = ", m)
		if x < slice[m] {
			l = m + 1
		} else {
			index = m
			fmt.Println("index = ", index)
			r = m - 1
		}
		m = (l + r) / 2
	}
	return index
}

func binarySearch2(slice []int, n, x int) int {
	l, r := 0, n-1
	m := (l + r) / 2

	index := -1
	for l <= r {
		if x <= slice[m] {
			index = m
			l = m + 1
		} else {
			r = m - 1
		}
		m = (l + r) / 2
	}
	return index
}
