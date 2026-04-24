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

		var str string
		fmt.Fscan(reader, &str)

		stringMap := make(map[rune]int)
		for _, char := range str {
			stringMap[char]++
		}

		var maxFreuency int
		for _, val := range stringMap {
			if val > maxFreuency {
				maxFreuency = val
			}
		}

		//check if the most frequency char lies at the last position of string
		lastCharFrequency := stringMap[rune(str[n-1])]
		if maxFreuency == lastCharFrequency {
			fmt.Fprintln(writer, n-maxFreuency)
		} else {
			fmt.Fprintln(writer, n-stringMap[rune(str[n-1])])
		}
	}
}
