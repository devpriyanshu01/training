package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//input number of testcases
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		//take array input
		slice := make([]int, 7)
		for j := 0; j < 7; j++ {
			fmt.Fscan(reader, &slice[j])
		}

		//core logic
		sort.Ints(slice)

		for j := 0; j < 6; j++ {
			slice[j] = -1 * slice[j]
		}

		sum := 0
		for j := 0; j < 7; j++ {
			sum += slice[j]
		}

		fmt.Fprintln(writer, sum)
	}

}
