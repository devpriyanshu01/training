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

	var q int 
	fmt.Fscan(reader, &q)

	queries := make([]int, q)
	max := math.MinInt
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)
		queries[i] = x	

		if x > max {
			max = x
		}
	}

	//prepare isPrime slice
	spf := make([]int, max+1)
	for i := range spf {
		spf[i] = i
	}

	for i := 2; i < max; i++ {
		if spf[i] == i {
			for j := i*i; j <= max; j += i {
				spf[j] = min(spf[j], i)
			}
		}
	}

	//now find prime factorization of each queries
	
	for _, x := range queries {
		myMap := make(map[int]int)
		var factors []int
		for x > 1 {
			pf := spf[x]	
			factors = append(factors, pf)
			myMap[pf]++
			x = x/pf
		}

		for _, val := range factors {
			_, exists := myMap[val]
			if exists {
				fmt.Fprintf(writer, "%d^%d ", val, myMap[val])
				delete(myMap, val)
			}
		}
		fmt.Fprintln(writer)
	}
	
}
