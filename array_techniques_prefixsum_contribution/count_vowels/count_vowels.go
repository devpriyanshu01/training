package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//size of array
	n := 0
	fmt.Fscan(reader, &n)

	//fill array
	var str string
	fmt.Fscan(reader, &str)

	vowel := make(map[byte]bool)
	vowel['a'] = true
	vowel['e'] = true
	vowel['i'] = true
	vowel['o'] = true
	vowel['u'] = true

	//create prefixsum array
	prefix := make([]int, n)
	vcount := 0
	for i := 0; i < n; i++ {
		_, ok := vowel[str[i]]
		if ok {
			vcount++
			prefix[i] = vcount
		} else {
			prefix[i] = vcount
		}
	}

	//get value of queries
	q := 0
	fmt.Fscan(reader, &q)

	writer := bufio.NewWriter(os.Stdout)
	ans := 0

	for i := 0; i < q; i++ {
		l, r := 0, 0
		fmt.Fscan(reader, &l)
		fmt.Fscan(reader, &r)

		if l == 1 {
			ans = prefix[r-1]
		} else {
			ans = prefix[r-1] - prefix[l-2]
		}
		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
