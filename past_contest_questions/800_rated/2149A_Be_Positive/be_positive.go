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

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		slice := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &slice[i])
		}

		//logic begins here...
		freq := make(map[int]int)
		for _, val := range slice {
			freq[val]++
		}

		// log.Println("freq map", freq)

		ans := 0
		//operations for 0
		zero, ok := freq[0]
		if ok {
			ans += zero
		}
		// log.Println("zero", zero)

		minueOne, ok := freq[-1]
		if ok {
			e := minueOne % 2
			ans += e * 2
		}
		// log.Println("after minusone", ans)
		fmt.Fprintln(writer, ans)
	}

}
