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

		//input slice
		arr := make([]int, n)
		for i := range arr {
			fmt.Fscan(reader, &arr[i])
		}

		var distinctCount int
		distinctCount++

		for i := 1; i < n; i++ {
			//check if arr[i] already exists
			var present = false
			for j := 0; j < i; j++ {
				if arr[j] == arr[i] {
					present = true
				}
			}
			if !present {
				distinctCount++
			}
		}

		//check check if distinctCount present in the arr or not.
		//if present that is the answer, else ans will be a number just greater than distinct number.
		//sort the arr
		sort.Ints(arr)

		for _, val := range arr {
			if val >= distinctCount {
				fmt.Fprintln(writer, val)
				break
			}
		}
	}

}
