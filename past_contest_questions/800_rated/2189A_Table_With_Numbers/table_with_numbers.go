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
		var n, h, l int
		fmt.Fscan(reader, &n, &h, &l)

		//array input
		nedNumbers := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &nedNumbers[i])
		}

		//logic is that:
		//idea is that - if the sum of pairs > sum of (h+l) then it will not form a valid index of a table.
		//So this pairs is not possible

		var ans int
		


				if nedNumbers[i] != -1 && nedNumbers[j] != -1 {
					if nedNumbers[i] + nedNumbers[j] <= sum {
						ans++
						arr[i] = -1
						arr[j] = -1
					}
				}
			}
		}