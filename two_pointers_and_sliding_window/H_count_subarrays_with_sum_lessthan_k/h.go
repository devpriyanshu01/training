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

	//n & k input
	var n, k int
	fmt.Fscan(reader, &n, &k)

	//slice input
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	//core logic begins...
	i := 0
	ans, sum := 0, 0

	for j := 0; j < n; j++ {
		sum += slice[j]

		for sum >= k {
			sum -= slice[i]
			i++
		}

		ans += j - i + 1
	}

	fmt.Fprintln(writer, ans)

}
