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
		var l, a, b int
		fmt.Fscan(reader, &l, &a, &b)

		maxProfit := a
		count := 0
		pos := a
		
		for count <= l {
			pos += b
			profit := pos % l
			if profit > maxProfit {
				maxProfit = profit
			}
			count++
		}
		fmt.Fprintln(writer, maxProfit)
	}


}
