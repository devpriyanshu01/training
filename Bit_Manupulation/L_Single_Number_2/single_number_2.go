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

	var n int
	fmt.Fscan(reader, &n)

	//slice input
	peoples := make([]int, n)
	for i := range peoples {
		fmt.Fscan(reader, &peoples[i])
	}

	//logic begins here...
	//Idea : Since only one person is present once & others are all in the multiples of 3.
	//The sum of bits at any posn. must be multiple of 3 without the present of imposter.
	//Only when imposter is present this multiple of three pattern is disturbed.

	imposter := 0
	for i := 0; i < 64; i++ {
		mask := (1 << i)
		bitsSum := 0
		for _, people := range peoples {
			if people&mask != 0 {
				bitsSum++
			}
		}
		if (bitsSum % 3) != 0 { //set this bit as imposter contributes 1 bit here.
			imposter = (imposter | mask)
		}
	}
	fmt.Fprintln(writer, imposter)
}
