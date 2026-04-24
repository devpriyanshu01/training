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
		fmt.Fscan(reader,&n)

		slice := make([]int, n)
		for i := range slice {
			fmt.Fscan(reader, &slice[i])
		}

		//core logic goes here...
		for i := 0; i < n; i++ {
			for j := i+1; j < n; j++ {
				if slice[j] < slice[i] {
					if checkParity(slice[i], slice[j]) {
						swap(slice, i, j) 
						j = i+1
					}
				}
			}
		}

		for _, val := range slice {
			fmt.Fprint(writer, val, " ")
		}	
		fmt.Fprintln(writer)
	}
}

func checkParity(a, b int) bool {
	if a & 1 == 1 && b & 1 == 0 {
		return true
	}else if a & 1 == 0 && b & 1 == 1 {
		return true
	}
	return false
}

func swap(slice []int, l, r int) {
	temp := slice[l] 
	slice[l] = slice[r]
	slice[r] = temp
}
