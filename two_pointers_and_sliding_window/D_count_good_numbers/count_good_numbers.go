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

	n, m, k := 0, 0, 0
	fmt.Fscan(reader, &n, &m, &k)

	//slice input
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	//good numbers
	goodNum := make(map[int]struct{})
	for i := 0; i < m; i++ {
		var gn int
		fmt.Fscan(reader, &gn)
		goodNum[gn] = struct{}{}
	}

	//core logic begins here...
	//build the window
	match := 0
	for i := 0; i < k; i++ {
		_, ok := goodNum[slice[i]]	
		if ok {
			match++
		}
	}
	fmt.Fprint(writer, match, " ")

	//slide the window
	for i := 0; i < (n-k); i++ {
		//remove last element of the current window
		_, ok := goodNum[slice[i]]
		if ok {
			match--
		}

		//add next number to current window
		_, ok = goodNum[slice[i+k]]
		if ok {
			match++
		}

		fmt.Fprint(writer, match, " ")
	}
}
