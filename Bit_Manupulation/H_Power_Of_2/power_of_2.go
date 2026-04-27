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

	var n uint64
	fmt.Fscan(reader, &n)

	var count uint64
	for i := 0; i < 64; i++ {
		var mask uint64 = 1 << i
		if (n & mask) != 0 {
			count++
		}
	}

	if count == 1 {
		fmt.Fprintln(writer, "YES")
	}else {
		fmt.Fprintln(writer, "NO")
	}
}