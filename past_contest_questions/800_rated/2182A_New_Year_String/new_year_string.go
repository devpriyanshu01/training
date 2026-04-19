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

		// log.Printf("for string %s", str)

		newYearString := false
		for j := 0; j < n-3; j++ {
			if str[j:j+4] == "2026" {
				//a new year string
				fmt.Fprintln(writer, 0)
				newYearString = true
				break
			}
		}

		//if newYearString == true => 2026 is found & hence given string is already a new year string.
		//so we don't do anything

		if !newYearString {
			// log.Printf("%s ;string does not contain 2026", str)
			for j := 0; j < n-3; j++ {
				if str[j:j+4] == "2025" {
					fmt.Fprintln(writer, 1)
					newYearString = true
					break
				}
			}
		}

		//even if at this point newYearString == false, implies 2025 string is not present in given string.
		//so given string is already a newYearString.
		if !newYearString {
			// log.Printf("%s ;string does not contain 2025", str)
			fmt.Fprintln(writer, 0)
		}
	}
}
