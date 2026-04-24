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

	//number of testcases
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, s, x int
		fmt.Fscan(reader, &n, &s, &x)

		//slice elements
		slice := make([]int, n)
		for  i := range slice {
			fmt.Fscan(reader, &slice[i])
		}

		sum := 0
		for _, val := range slice {
			sum += val
		}

		if sum == s {
			fmt.Fprintln(writer, "Yes")
		}else if sum > s {
			fmt.Fprintln(writer, "No")
		}else {
			toAdd := s - sum
			if toAdd % x == 0 {
				fmt.Fprintln(writer, "Yes") //when possible to add
			}else {
				fmt.Fprintln(writer, "No")
			}
		}
	}
}
