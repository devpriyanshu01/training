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

	for i := 0; i < q; i++ {
		var n int
		fmt.Fscan(reader, &n)

		//input s & t
		var s, t string
		fmt.Fscan(reader, &s, &t)

		sMap := make(map[rune]int)
		tMap := make(map[rune]int)

		for _, char := range s {
			sMap[char]++
		}

		for _, char := range t {
			tMap[char]++
		}

		//check if they have the same runes with same frequency
		misMatch := false
		for key, val := range sMap {
			count, ok := tMap[key]
			if ok {
				if val != count {
					fmt.Fprintln(writer, "NO")
					misMatch = true
					break
				}
			} else {
				fmt.Fprintln(writer, "NO")
				misMatch = true
				break
			}
		}

		//
		if !misMatch {
			fmt.Fprintln(writer, "YES")
		}
	}
}
