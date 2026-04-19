package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//size of array
	n := 0
	fmt.Fscan(reader, &n)

	//fill in the array
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &array[i])
	}

	//generate prefix sum array
	prefix := make([]int, n)
	sum := 0
	for i := 1; i <= n; i++ {
		if i&1 == 1 {
			//index is odd
			prefix[i-1] = sum
		} else {
			//index is even
			sum += array[i-1]
			prefix[i-1] = sum
		}
	}

	//logic begins from here...
	//get no. of queries
	q := 0
	fmt.Fscan(reader, &q)

	ans := 0
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < q; i++ {
		l, r := 0, 0
		fmt.Fscan(reader, &l)
		fmt.Fscan(reader, &r)

		//when l is 1
		if l == 1 {
			ans = prefix[r-1]
		} else {
			ans = prefix[r-1] - prefix[l-2]
		}
		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
