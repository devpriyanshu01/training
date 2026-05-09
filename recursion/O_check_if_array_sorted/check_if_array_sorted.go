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

    isSorted := f(slice, t)
    if isSorted {
    	fmt.Fprintln(writer, "Yes")
    }else {
    	fmt.Fprintln(writer, "No")
    }
}

func f(slice []int, n int) bool {
	if n == 1 {
		return true
	}

	return f(slice, n-1) && slice[n-1] >= slice[n-2]
}