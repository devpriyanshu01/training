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
		var n, k int
		fmt.Fscan(reader, &n, &k)

		//input string
		var str string
		fmt.Fscan(reader, &str)

		var totalSleep, i, cantSleep int

		for i < n {
			if str[i] == '1' {
				cantSleep = k
			} else {
				if cantSleep == 0 {
					totalSleep++
				}else {
					cantSleep--
				}
			}
			i++
		}
		fmt.Fprintln(writer, totalSleep)
	}
}
