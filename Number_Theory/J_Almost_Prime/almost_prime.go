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

	var n int
	fmt.Fscan(reader, &n)

	spf := make([]int, n+1)
	//initially we mark all the numbers as prime by making spf[i] = i
	for i := range spf {
		spf[i] = i
	}

	//Property of prime number : the smallest prime factor will be the number itself.
	//Apply this property
	for i := 2; i <= n; i++ {
		if spf[i] == i { //meaning its a prime number
			for j := i * i; j <= n; j += i {
				spf[j] = min(spf[j], i)
			}
		}
	}

	var ans int
	for i := 2; i <= n; i++ {
		set := make(map[int]struct{})
		x := i
		for x > 1 {
			pf := spf[x]
			set[pf] = struct{}{}
			x /= pf
		}
		if len(set) == 2 {
			ans++
		}
	}

	fmt.Fprintln(writer, ans)

}
