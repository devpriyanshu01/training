package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		//logic begins here...
		initOne := countRune(str, '1')

		var sb strings.Builder
		for i := 1; i < n-1; i++ {
			if str[i-1] == '1' && str[i+1] == '1' {
				//we can select this index & change it to one
				sb.WriteString(str[:i])
				sb.WriteRune('1')
				sb.WriteString(str[i+1:])

				str = sb.String()
				sb.Reset()
			}
		}

		maxOne := countRune(str, '1')

		//find minimum count of 1 possible
		sb.Reset()
		for i := 1; i < n-1; i++ {
			if str[i-1] == '1' && str[i+1] == '1' {
				sb.WriteString(str[:i])
				sb.WriteRune('0')
				sb.WriteString(str[i+1:])

				str = sb.String()
				sb.Reset()
			}
		}

		//count 1's now
		minOne := countRune(str, '1')
		minOne = min(minOne, initOne)
		fmt.Fprintln(writer, minOne, maxOne)
	}
}

func countRune(str string, char byte) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			count++
		}
	}
	return count
}
