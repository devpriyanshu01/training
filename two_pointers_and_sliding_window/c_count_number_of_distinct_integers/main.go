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

	n, k := 0, 0
	fmt.Fscan(reader, &n, &k)

	//insert array elements
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	//main logic
	//build the window
	freq := make(map[int]int)
	for i := 0; i < k; i++ {
		freq[slice[i]]++
	}

	fmt.Fprint(writer, len(freq), " ")

	//slice the window
	for i := 0; i < (n - k); i++ {
		//remove the last element
		freq[slice[i]]--
		val, _ := freq[slice[i]]
		if val == 0 {
			delete(freq, slice[i])
		}

		//add next element
		freq[slice[i+k]]++

		fmt.Fprint(writer, len(freq), " ")
	}
}
