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

		permutation := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &permutation[i])
		}

		//logic begins here...
		permuMap := make(map[int]int)
		for i, val := range permutation {
			permuMap[val] = i
		}

		//create a copy of the array for sorting
		copyPer := make([]int, n)
		// for i, val := range permutation {	//try using copy function here
		// 	copyPer[i] = val
		// }

		//using copying
		copy(copyPer, permutation)

		sort.Ints(copyPer) //copy the slice

		var l, r int
		shouldBe := 0
		for i := n - 1; i >= 0; i-- {
			val := permuMap[copyPer[i]]
			// fmt.Printf("%d is at index := %d;;;; should be at %d", copyPer[i], val, shouldBe)
			if val != shouldBe {
				r = val
				l = shouldBe
				break
			} else {
				shouldBe++
			}
		}

		//now we need to reverse the given permuted array
		reverseSlice(permutation, l, r)

		for i := 0; i < n; i++ {
			fmt.Fprint(writer, permutation[i], " ")
		}
		fmt.Fprintln(writer)
	}
}

func reverseSlice(slice []int, l, r int) {
	size := (r - l + 1)
	if size&1 == 1 {
		size = (size / 2) + 1
	} else {
		size = size / 2
	}

	for i := 0; i < size; i++ {
		temp := slice[l]
		slice[l] = slice[r]
		slice[r] = temp
		r--
		l++
	}
}
