package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n uint32
	fmt.Fscan(reader, &n)

	var count uint32
	for i := 0; i < 32; i++ {
		var mask uint32 = 1 << i
		if (n & mask) != 0 {
			count++
		}
	}

	fmt.Fprintln(writer, count)

}