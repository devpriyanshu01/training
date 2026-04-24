package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//take input of n & k
	n, k := 0, 0
	fmt.Fscan(reader, &n, &k)

	//input array elements
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	//core logic
	max, sum := math.MinInt, 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if sum > max {
		max = sum
	}

	for i := 0; i < (n-k); i++ {
		sum -= arr[i]
		sum += arr[k+i]
		if sum > max {
			max = sum
		}
	}
	fmt.Fprintln(writer, max)
}
