package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	
	n := 0
	fmt.Fscan(in, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	count := make(map[int]int)
	for _, v := range arr {
		count[v] = count[v]+1
	}
	
	  fmt.Println(len(count))
}
