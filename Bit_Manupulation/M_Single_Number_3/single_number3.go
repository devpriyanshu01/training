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

	var n int
	fmt.Fscan(reader, &n)

	peoples := make([]int, n)
	for i := range peoples {
		fmt.Fscan(reader, &peoples[i])
	}

	//core logic begins here...

	//Idea : all present twice except the 2 imposter.
	//If take take xor of all the numbers then we get - xor of imposter1 and imposter2
	//Now, we find at what posn they have set bit => meaning imposters are different at that posn.
	//Divide the people into 2 groups on the basis of that posn bits.
	//then take xor of both groups.
	//imposter1 & imposter 2 will be revealed

	xor := peoples[0]
	for i, people := range peoples {
		if i == 0 {
			continue
		}
		xor ^= people
	}

	//find at what posn xor has set bit
	posn := 0
	for i := range 64 {
		mask := 1 << i
		if xor&mask != 0 {
			posn = i
		}
	}

	//time to separate in groups
	var group1, group2 []int

	mask := 1 << posn
	for _, people := range peoples {
		if people&mask == 0 {
			group1 = append(group1, people)
		} else {
			group2 = append(group2, people)
		}
	}

	//check for first imposter in group1
	imposter1 := group1[0]
	for i, people := range group1 {
		if i == 0 {
			continue
		}
		imposter1 ^= people
	}

	//look for the second imposter in group2
	imposter2 := group2[0]
	for i, people := range group2 {
		if i == 0 {
			continue
		}
		imposter2 ^= people
	}

	if imposter1 < imposter2 {
		fmt.Fprintln(writer, imposter1, imposter2)
	} else {
		fmt.Fprintln(writer, imposter2, imposter1)
	}
}
