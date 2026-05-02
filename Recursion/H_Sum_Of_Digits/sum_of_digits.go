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

	var n int
	fmt.Fscan(reader, &n)

	fmt.Fprintln(writer, f(n))

}

func f(n int) int {
	if n == 0 {
		return 0
	}

	return (n % 10) + f(n/10)
}