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

        //logic starts here...
        for i := n-2; i >= 0; i-- {
        	if slice[i] >= 0 && slice[i+1] > 0 {
        		slice[i] += slice[i+1]
         	}else if slice[i] < 0 && slice[i+1] > slice[i] {
          		slice[i] += slice[i+1]
          }
        }

        // for i := n-2; i >= 0; i-- {
        // 	if slice[i] == 0 && slice[i+1] > 0 {
        //  		slice[i] += slice[i+1]
        //  }
        // }

        var ans int
        for _, num := range slice {
        	if num > 0 {
         		ans++
         }
        }
       fmt.Fprintln(writer, ans)
    }
}