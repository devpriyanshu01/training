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

	// Read one character (rune)
	char, _, err := reader.ReadRune()
	if err != nil {
		return
	}

	// Logic to categorize the character
	if char >= 'a' && char <= 'z' {
		fmt.Fprintln(writer, "Lowercase")
	} else if char >= 'A' && char <= 'Z' {
		fmt.Fprintln(writer, "Uppercase")
	} else if char >= '0' && char <= '9' {
		fmt.Fprintln(writer, "Digit")
	} else {
		fmt.Fprintln(writer, "Special")
	}
}