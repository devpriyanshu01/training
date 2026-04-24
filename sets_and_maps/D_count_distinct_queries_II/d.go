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

	//q
	var q int
	fmt.Fscan(reader, &q)

	multiSet := make(map[int]int)

	for i := 0; i < q; i++ {
		var queryType, x int
		fmt.Fscan(reader, &queryType)
		
		if queryType != 3 {
			fmt.Fscan(reader, &x)
		}

		switch queryType {
		case 1:
			multiSet[x]++
		case 2:
			//check if x exists in the map
			_, ok := multiSet[x]
			if ok {
				multiSet[x]--
				val := multiSet[x]
				if val == 0 {
					delete(multiSet, x)
				}
			}
		case 3:
			fmt.Fprintln(writer, len(multiSet))
		case 4:
			_, ok := multiSet[x]
			if ok {
				fmt.Fprintln(writer, "YES")
			}else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}
}
