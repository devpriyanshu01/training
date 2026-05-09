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

        //print odd values
        var till, start int
        if n & 1 == 1 {
        	till = n
        	start = n-1
        }else {
        	till = n-1
         	start = n
        }

        for i := 1; i <= till; i += 2 {
        	fmt.Fprint(writer, i, " ")
        }

        //for even nos.
        for i := start; i >= 2; i -= 2 {
        	fmt.Fprint(writer, i, " ")
        }

        fmt.Fprintln(writer)
    }
}