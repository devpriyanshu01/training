package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(reader, &nums[i])
	}

	//logic begins here...
	for i := 1; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			mask := 1 << j
			if i & mask != 0 {
				fmt.Fprint(writer, nums[j], " ")
			}
		}
		fmt.Fprintln(writer)
	}
}