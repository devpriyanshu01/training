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
        var till int
        if n & 1 == 1 {
        	till = n-2
        }else {
        	till = n-1
        }

        for i := 1; i <= till; i += 2 {
        	//swap
         	temp := slice[i-1]
          	slice[i-1] = slice[i]
           	slice[i] = temp
        }

        //print the answer.
        for _, val := range slice {
        	fmt.Fprint(writer, val, " ")
        }
        fmt.Fprintln(writer)
    }
}