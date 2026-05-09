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

    slice := make([]int, t)
    for i := 0; i < t; i++ {
    	fmt.Fscan(reader, &slice[i])
    }

    sum := f(slice, t)
    fmt.Fprintln(writer, sum)

}

func f(slice []int, n int) int {
	if n == 1 {
		return slice[n-1]
	}

	sum := f(slice, n-1)
	return sum + slice[n-1]
}