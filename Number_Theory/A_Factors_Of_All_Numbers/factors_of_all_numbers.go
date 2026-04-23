package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	// defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	//create n+1 sized 2d array
	factors := make([][]int, n+1)

	for i := 1; i <= n; i++ {

		for j := i; j <= n; j = j + i {
			factors[j] = append(factors[j], i)
		}

	}

	//writer insertion
	for i := 0; i < len(factors); i++ {
		if i == 0 {
			continue
		} else {
			for _, val := range factors[i] {
				fmt.Fprint(writer, val, " ")
			}
			fmt.Fprintln(writer)
		}
	}

	end := time.Since(start)
	writer.Flush()
	fmt.Println("Time Taken:", end)
}

/*
func factorsOfN(n int, writer io.Writer) {
	var last []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Fprint(writer, i, " ")

			if i != n/i {
				last = append(last, n/i)
			}
		}
	}

	for i := len(last)-1; i >= 0; i-- {
		fmt.Fprint(writer, last[i], " ")
	}

}
*/
