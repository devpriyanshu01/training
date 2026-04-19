package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//get the value of n, q, k
	n, q, k := 0, 0, 0
	fmt.Fscan(reader, &n, &q, &k)

	//fill array with elements
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &array[i])
	}

	//create prefix sum array
	digitSumK := 0
	prefix := make([]int, n)
	for i := 0; i < n; i++ {
		curr := array[i]
		sum := 0
		for curr > 0 {
			r := curr % 10
			curr = curr / 10
			sum += r
		}
		if sum == k {
			prefix[i] = digitSumK + 1
			digitSumK++
		} else {
			prefix[i] = digitSumK
		}
	}

	//create a writer
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//core logic begins
	ans := 0
	for i := 0; i < q; i++ {
		l, r := 0, 0
		fmt.Fscan(reader, &l, &r)

		if l == 1 {
			ans = prefix[r-1]
		} else {
			ans = prefix[r-1] - prefix[l-2]
		}
		fmt.Fprintln(writer, ans)
	}
}
