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

	//logic begins here...
	//create a powers of 2 slice

	p := 0
	var powers []int

	currPower := int(math.Pow(2, float64(p)))
	for currPower <= n {
		powers = append(powers, currPower)
		p++
		currPower = int(math.Pow(2, float64(p)))
	}

	i := len(powers)-1
	for n > 0 {
		if powers[i] <= n {
			fmt.Fprint(writer, powers[i], " ")
			n -= powers[i]
		}
		i--
	}
	fmt.Fprintln(writer)
}
