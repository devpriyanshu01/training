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
        taken := make([]bool, n)
        //modify the positive ones
        for i := 0; i < n-1; i++ {
        	if slice[i] > 0 && slice[i+1] > 0 {
         		slice[i] += slice[i+1]
           		taken[i] = true
         	} 
        }
        // fmt.Println("after working on positive numbers", slice)

        //now we will work on -ve numbers
        for i := 0; i < n-1; i++ {
        	if slice[i] < 0 {
         		slice[i] += slice[i+1]
           		taken[i] = true
         }
        }
        
        // fmt.Println("after working on negative numbers", slice)

        //now work on 0 in the slice
        for i := 0; i < n-1; i++ {
        	if slice[i] == 0 {
         		slice[i] += slice[i+1]
         }
        }
        // fmt.Println("after working on zeros numbers", slice)
       
       //now count how many positive integers are present in tthe lsice
       ans := 0
       for _, num := range slice {
           if num > 0 {
               ans++
           }
       } 
       fmt.Fprintln(writer, ans)
    }
}