package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//testcases
	var t int
	fmt.Fscan(reader, &t)

	//allowed moves
	var allowedMoves = [][]int{{2, 1}, {3, 0}, {4, -1}}

	for i := 0; i < t; i++ {
		var destX, destY int
		fmt.Fscan(reader, &destX, &destY)

		initX, initY := 0, 0

		success := false
		
		solve(allowedMoves, initX, initY, destX, destY, writer, &success)
		if !success {
			fmt.Fprintln(writer, "No")
		}
		
	}
}

func solve(allowedMoves [][]int, initX int, initY int, destX int, destY int, writer io.Writer, success *bool) {
	if initX > destX || initY > destY {
		return
	}

	if initX == destX && initY == destY {
		*success = true
		fmt.Fprintln(writer, "Yes")
		return
	}

	log.Printf("(initX, initY) - (%d, %d)", initX, initY)

	for _, slice := range allowedMoves {
		initX += slice[0]
		initY += slice[1]
		solve(allowedMoves, initX, initY, destX, destY, writer, success)
	}
}
