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

	t := 0
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		n, q := 0, 0
		fmt.Fscan(reader, &n, &q)

		//input slice elements
		slice := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &slice[j])
		}

		for j := 0; j < q; j++ {
			decider, hike := 0, 0
			fmt.Fscan(reader, &decider, &hike)

			even := false
			if decider == 0 {
				even = true
			}

			//core logic begins here...
			for k := 0; k < n; k++ {
				if slice[k]&1 == 1 && !even {
					slice[k] += hike
				} else if slice[k]&1 == 0 && even {
					slice[k] += hike
				}
			}

			//calculate sum now for the modified slice
			sum := 0
			for _, val := range slice {
				sum += val
			}
			fmt.Fprintln(writer, sum)
		}
	}
}
