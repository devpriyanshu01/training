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
	
	//n & k
	var n, k int
	fmt.Fscan(reader, &n, &k)

	//slice input
	slice := make([]int, n)
	for i := range slice {
		fmt.Fscan(reader, &slice[i])
	}

	//core logic begins...
	//prepare 2 maps to help us to search quickly
	freq := make(map[int]int)
	checkIndex := make(map[int]int)

	//create freq map for handling duplicate elements
	for _, val := range slice {
		freq[val]++
	}

	//create checkIndex map to check the index of present elements
	for i, val := range slice {
		checkIndex[val] = i
	}

	//variable to know the case when there no two-sum exists
	found := false
	//main logic
	for i, val := range slice {
		reqd := k - val

		count, exist := freq[reqd]
		if exist {
			//handle duplicate elements
			if reqd == val {
				//in this case count should be minimum 2 eg. val = 6 then for x = 12, reqd = 6
				if count > 1 {
					fmt.Fprintln(writer, i+1, checkIndex[reqd]+1)
					found = true
					break
				}
			}else if reqd != val {
				fmt.Fprintln(writer, i+1, checkIndex[reqd]+1)
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Fprintln(writer, -1)
	}
}
