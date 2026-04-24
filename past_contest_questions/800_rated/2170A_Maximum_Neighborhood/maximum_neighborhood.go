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

		//create 2-d slice
		grid := make([][]int, n)
		for i := 0; i < len(grid); i++ {
			grid[i] = make([]int, n)
		}

		//fill the grid with values
		num := 1
		for i := 0; i < n; i++ {
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j] = num
				num++
			}
		}

		//calculate neighborhood cost
		cost := make([][]int, n)
		for i := 0; i < len(cost); i++ {
			cost[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < len(grid[i]); j++ {
				c := calcSum(grid, n, grid[i][j], i, j)
				cost[i][j] = c
			}
		}

		//calculate max cost
		maxCost := 0
		for i := 0; i < n; i++ {
			for j := 0; j < len(cost[i]); j++ {
				if cost[i][j] > maxCost {
					maxCost = cost[i][j]
				}
			}
		}

		//send the maxCost to buffer for printing
		fmt.Fprintln(writer, maxCost)
	}
}

func calcSum(grid [][]int, n, val, r, c int) int {
	sum := val

	//row -
	if r-1 >= 0 {
		sum += grid[r-1][c]
	}

	//row +
	if r+1 < n {
		sum += grid[r+1][c]
	}

	//column -
	if c-1 >= 0 {
		sum += grid[r][c-1]
	}

	//column +
	if c+1 < n {
		sum += grid[r][c+1]
	}
	
	return sum
}
