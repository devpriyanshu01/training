package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := 0
	fmt.Fscan(reader, &n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &array[i])
	}


	//reverse
	writer := bufio.NewWriter(os.Stdout)
	for i := n-1; i >= 0; i--{
		fmt.Fprint(writer, array[i], " ")
	}
	defer writer.Flush()
}
