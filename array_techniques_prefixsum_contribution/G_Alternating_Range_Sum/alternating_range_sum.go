package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Fscan(reader, &numbers[i])
	}

	//prefix array
	prefix := make([]int, n)
	polarity := true
	sum := 0
	for i, num := range numbers {
		if polarity {
			sum += num
			prefix[i] = sum
			polarity = !polarity
		}else {
			sum -= num
			prefix[i] = sum
			polarity = !polarity
		}
	}
	
	//start answering queries
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)

		l--
		r--

		var ans int
		if l == 0 {
			ans = prefix[r]
		}else {
			if l & 1 == 1 {
				//odd value of l
				ans = -1 * (prefix[r] - prefix[l-1])
			}else {
				//even value of l
				ans = prefix[r] - prefix[l-1]
			}
		}
		fmt.Fprintln(writer, ans)
	}

}