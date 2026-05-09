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

        b := make([]int, n)
        for j := range b {
        	fmt.Fscan(reader, &b[j])
        }

        //logic starts
        set := make(map[int]struct{})
        for _, val := range b {
        	set[val] = struct{}{}
        }

        fmt.Fprintln(writer, len(set))
    }
}