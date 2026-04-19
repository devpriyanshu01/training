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

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		//logic begins...
		var left, right int
		if n&1 == 1 {
			//odd number of elements
			left = n/2 + 1
			right = n / 2
		} else {
			//even number of elements
			left = n / 2
			right = n / 2
		}

		//slice to store the permutation that satisfies given condition
		permutation := make([]int, n)
		//filling elements from left to right
		num := 1
		j := n - 1
		for i := 0; i < left; i++ {
			permutation[j] = num
			num++
			j -= 2
		}

		//now start filling elements in permutation slice from right to left.
		if n&1 == 1 {
			//odd no. of elements
			j = 1
		} else {
			//even no. of elements
			j = 0
		}

		for i := 0; i < right; i++ {
			permutation[j] = num
			num++
			j += 2
		}

		//prints the elements of the resultant permuation in writer buffer
		for _, val := range permutation {
			fmt.Fprint(writer, val, " ")
		}
		fmt.Fprintln(writer)
	}
}
