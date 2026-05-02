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

	var n, q, k int
	fmt.Fscan(reader, &n, &q, &k)

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Fscan(reader, &numbers[i])
	}

	//store the count of each numbers[i] in factors array
	factors := make([]int, n)
	for i, num := range numbers {
		count := findFactorsCount(num)
		factors[i] = count
	}

	//find the prefix array
	prefix := make([]int, n)
	kFact := 0
	for i, fact := range factors {
		if fact == k {
			prefix[i] = kFact + 1
			kFact++
		}else {
			prefix[i] = kFact
		}
	}

	//start answering the query
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		l--
		r--

		var ans int
		if l == 0 {
			ans = prefix[r]
		}else {
			ans = prefix[r] - prefix[l-1]
		}	
		fmt.Fprintln(writer, ans)
	}
}

func findFactorsCount(num int) int {
	count := 0

	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			if i == num/i {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}
