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

		var a, b string
		fmt.Fscan(reader, &a, &b)

		isAValid := checkValidSequence(a)
		isBValid := checkValidSequence(b)
		if isAValid && isBValid {
			fmt.Fprintln(writer, "Yes")
			// fmt.Println("both a & b were already valid.")
			continue
		}

		//try to make a & b regular.
		if a[0] == ')' || b[0] == ')' {
			fmt.Fprintln(writer, "No")
			// fmt.Println("not possible in first check.")
			continue
		}
		// open := 1
		solved := false
		var asb, bsb strings.Builder
		for i := 1; i < n; i++ {
			if a[i] != b[i] {
				asb.WriteString(a[:i])
				asb.WriteRune(rune(b[i]))
				asb.WriteString(a[i+1:])

				//swap b
				bsb.WriteString(b[:i])
				bsb.WriteRune(rune(a[i]))
				bsb.WriteString(b[i+1:])

				//check this change make both string regular
				isARegular := checkValidSequence(asb.String())
				isBRegular := checkValidSequence(bsb.String())
				if isARegular && isBRegular {
					fmt.Fprintln(writer, "Yes")
					// fmt.Printf("swapping at i = %d, makes regular\n", i)
					solved = true
					break
				}
			}
		}

		if !solved {
			fmt.Fprintln(writer, "No")
			// fmt.Println("not possible make them regular , no hope")
			continue
		}

	}
}

func checkValidSequence(str string) bool {
	open := 0

	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			open++
		} else if str[i] == ')' && open == 0 {
			return false
		} else if str[i] == ')' && open > 0 {
			open--
		}
	}
	if open > 0 {
		return false
	}
	return true
}
