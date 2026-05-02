package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)

		//logic starts here...
		currX, currY := 0, 0

		solve(currX, currY, x, y, writer, true)

	}

}

func solve(cx, cy, x, y int, writer io.Writer, short bool) {
	//check if equal
	if cx == x && cy == y {
		fmt.Fprintln(writer, "YES")
		return
	}

	//if greater return
	if cx > x && cy > y {
		return
	}

	cx += 2
	solve(cx, cy, x, y, writer, short)

	cy += 2 
	solve(cx, cy, x, y, writer, short)

	if short && x-cx == 1{
		cx += 1
		short = false
		solve(cx, cy, x, y, writer, short)
	}

	if short && y-cy == 1 {
		cy += 1
		short = false
		solve(cx, cy, x, y, writer, short)
	}
	
}
