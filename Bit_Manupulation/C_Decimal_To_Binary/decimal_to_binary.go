package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	if n == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	//logic begins here
	p := 0
	currPower := int(math.Pow(2, float64(p)))
	var powers []int

	//create the array that has the powers of <= n
	for currPower <= n {
		powers = append(powers, currPower)
		p++
		currPower = int(math.Pow(2, float64(p)))
	}

	i := len(powers) - 1
	for i >= 0 {
		if powers[i] <= n {
			fmt.Fprint(writer, 1)
			n -= powers[i]
		}else {
			fmt.Fprint(writer, 0)
		}
		i--
	}
	fmt.Fprintln(writer)

}
