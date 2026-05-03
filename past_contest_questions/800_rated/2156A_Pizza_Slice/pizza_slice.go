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

		ans := f(n)
		fmt.Fprintln(writer, ans)
	}

}

//fn. with logic to solve this problem.
func f(n int) int {
	if n <= 2 {
		return 0
	}

	m1 := n/3
	var m2, m3 int
	left := n - m1

	//divide among m2 & m3.
	m2 = left/2
	m3 = left-m2

	fmt.Printf("for n: %d, m1 = %d, m2 = %d, m3 = %d\n", n, m1, m2, m3)
	
	return m1 + f(m3)
}
