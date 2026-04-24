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

	//number of testcases
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)

		//logic begins here...
		//Case 1. when top layer is choosen is white chocolate

		ansWhite, reqd := 0, 1
		choose := true
		availableWhite, availableDark := a, b

		for {
			//choosing white chocolate logic - if true, choose white else choose dark chocolate.
			if choose {
				if availableWhite >= reqd {
					ansWhite++
					availableWhite -= reqd
					reqd *= 2
					choose = false
				} else {
					break
				}
			} else {
				//here dark chocolate is choosen, when choose variable is false
				if availableDark >= reqd {
					ansWhite++
					availableDark -= reqd
					reqd *= 2
					choose = true
				} else {
					break
				}
			}
		}

		//case 2. when dark chocolate is choose to be in top layer
		ansDark, reqd := 0, 1
		availableWhite, availableDark = a, b
		choose = true

		for {
			//when true choose dark chocolate else choose white chocolate
			if choose {
				if availableDark >= reqd {
					ansDark++
					availableDark -= reqd
					reqd *= 2
					choose = false
				} else {
					break
				}
			} else {
				//white chocolate is choosen here
				if availableWhite >= reqd {
					ansDark++
					availableWhite -= reqd
					reqd *= 2
					choose = true
				} else {
					break
				}
			}
		}

		if ansWhite >= ansDark {
			fmt.Fprintln(writer, ansWhite)
		} else {
			fmt.Fprintln(writer, ansDark)
		}
	}
}
