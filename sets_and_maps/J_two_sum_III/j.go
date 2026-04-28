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

	//n & k
	var n, k int
	fmt.Fscan(reader, &n, &k)

	//slice input
	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	//store the value of the given slice into a map with their indexes as values.
	myMap := make(map[int]int)
	count := 0

	for _, val := range slice {
		freq, exists := myMap[k-val]
		if exists {
			count += freq
			// delete(myMap, k-val)
		}
		myMap[val]++
	}

	fmt.Fprintln(writer, count)
}
