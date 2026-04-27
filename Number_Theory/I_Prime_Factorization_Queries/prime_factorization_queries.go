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

	var q int
	fmt.Fscan(reader, &q)

	max := 1000000

	//prepare isPrime slice
	spf := make([]int, max+1)
	for i := range spf {
		spf[i] = i
	}

	for i := 2; i < max; i++ { //logic is crucial
		if spf[i] == i {
			for j := i * i; j <= max; j += i {
				spf[j] = min(spf[j], i)
			}
		}
	}

	//now find prime factorization of each queries

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)

		for x > 1 {
			pf := spf[x]
			e := 0
			for x % pf == 0 {
				x = x/pf
				e++
			}
			fmt.Fprintf(writer, "%d^%d ", pf, e)
		}
		fmt.Fprintln(writer)
	}

}
