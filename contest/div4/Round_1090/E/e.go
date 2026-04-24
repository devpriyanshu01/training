package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//input number of testcases
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		//array size
		var n int
		fmt.Fscan(reader, &n)

		//array elements
		slice := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &slice[j])
		}

		for j := 1; j < n; j++ { //no. of test cases
			//step 1. select
			s := len(slice)-1
			selectedNum := slice[s]

			//step 2 - xor all elements with selected elements
			k := 0
			for k < len(slice)-1 {
				slice[k] = slice[k] ^ selectedNum 
				k++
			}

			//step 3 - remove selected elements
			slice = slice[:len(slice)-1]
			log.Println("Remaining slice", slice)
		}
		fmt.Fprintln(writer, slice[0])
	}


}