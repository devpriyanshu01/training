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

        slice := make([]int, n)
        for i := range slice {
        	fmt.Fscan(reader, &slice[i])
        }

        //logic begins here...
        var ans = slice[0]
        for i := 1; i < n; i++ {
        	ans ^= slice[i]
        }

        fmt.Fprintln(writer, ans)
    }

}