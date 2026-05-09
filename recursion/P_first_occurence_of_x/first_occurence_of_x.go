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

    var x int
    fmt.Fscan(reader, &x)

    ans := f(slice, t, x)
    if ans == -1 {
    	fmt.Fprintln(writer, ans)
    }else {
    	fmt.Fprintln(writer, ans+1)
    }
}

func f(slice []int, n, x int) int{
	if n == 0 {
		return -1
	}

	index := f(slice, n-1, x)

	if index != -1 {
		return index
	}

	if slice[n-1] == x {
		return n-1
	}
	return -1
}