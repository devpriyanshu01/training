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

	//n & k
	var n int
	var k int64
	fmt.Fscan(reader, &n, &k)

	//slice input
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	//core logic begins...
	i, ans, sum := 0, math.MaxInt, 0
	for j := 0; j < n; j++ {
		sum += slice[j]

		//when sum is greater than k - we need to shrink
		for sum > int(k) {
			ans = min(ans, j-i+1)
			sum -= slice[i]
			i++
		}
	}

	if ans == math.MaxInt {
		fmt.Fprintln(writer, -1)
	}else {
		fmt.Fprintln(writer, ans)
	}
}


