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

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Fscan(reader, &numbers[i])
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := i; k <= j; k++ {
				fmt.Fprint(writer, numbers[k], " ")
			}
			fmt.Fprintln(writer)
		}
	}

}