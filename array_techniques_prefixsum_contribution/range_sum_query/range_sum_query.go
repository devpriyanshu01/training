package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	n := 0	//length of array
	fmt.Fscan(reader, &n)

	//take input of array
	arr := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &arr[i])
	}

	//prefix sum array
	prefixSum := make([]int, n)
	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}

	//number of query input
	q := 0
	fmt.Fscan(reader, &q)

	//create a writer to print result on terminal
	writer := bufio.NewWriter(os.Stdout)
	for range q {
		l, r := 0, 0
		fmt.Fscan(reader, &l)
		fmt.Fscan(reader, &r)

		ans := 0
		if l == 1 {
			ans = prefixSum[r-1]
		}else {
			ans = prefixSum[r-1] - prefixSum[l-2]
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}