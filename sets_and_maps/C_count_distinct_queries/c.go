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

	//q number
	var q int
	fmt.Fscan(reader, &q)

	//create multiset(slice)
	multiset := make(map[int]int)
	for i := 0; i < q; i++ {
		var queryType, x int
		fmt.Fscan(reader, &queryType)
		if queryType == 3 {
			x = 0
		} else {
			fmt.Fscan(reader, &x)
		}

		switch queryType {
		case 1:
			multiset[x]++
			// fmt.Println("ADDED", x)
		case 2:
			delete(multiset, x)
			// fmt.Println("DELETED")
		case 3:
			fmt.Fprintln(writer, len(multiset))
		case 4:
			_, ok := multiset[x]
			if ok {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}
}
