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

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		rect := make([][]int, 2)

		var n int
		fmt.Fscan(reader, &n)

		for i := 0; i < 2; i++ {
			rect[i] = make([]int, n)
		}

		//take color pattern input
		var first, second string
		fmt.Fscan(reader, &first, &second)

		for i := 0; i < n; i++ {
			rect[0][i] = int(first[i])
		}

		for i := 0; i < n; i++ {
			rect[1][i] = int(second[i])
		}

		log.Println("Given Input", rect)

		//logic begins here.
		pairs, operations := 0, 0
		for i := 0; i < 1; i++ {
			for j := 0; j < n; j++ {
				//check adjacent pair in first line
				if rect[i][j] == 'D' {
					continue
				}else if j == n-1 {
					if
				}else {
					if rect[i][j] == rect[i][j+1] {
						rect[i][j] = 'D'
						rect[i][j+1] = 'D'
						pairs++
					} else if rect[i][j] == rect[i+1][j] {
						rect[i][j] = 'D'
						rect[i+1][j] = 'D'
						pairs++
					} else {
						operations++
						pairs++
						rect[i][j] = 'D'
						rect[i][j+1] = 'D'
					}
				}

			}
		}
	}
}
