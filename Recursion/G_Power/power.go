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

	var x, n int
	fmt.Fscan(reader, &x, &n)

	fmt.Fprintln(writer, f(x, n))	

}

func f(x, n int) int {
	if n == 0 {
		return 1
	}

	return x * f(x, n-1)
}