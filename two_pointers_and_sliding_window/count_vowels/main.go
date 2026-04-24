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

	//take input n, k
	n, k := 0, 0
	fmt.Fscan(reader, &n, &k)

	//take input of string
	var str string
	fmt.Fscan(reader, &str)

	//create a map of all the vowels
	vowel := make(map[byte]bool)
	vowel['a'] = true
	vowel['e'] = true
	vowel['i'] = true
	vowel['o'] = true
	vowel['u'] = true
	vowel['A'] = true
	vowel['E'] = true
	vowel['I'] = true
	vowel['O'] = true
	vowel['U'] = true

	//build the window
	count := 0
	for i := 0; i < k; i++ {
		_, ok := vowel[str[i]]
		if ok {
			count++
		}
	}
	fmt.Fprint(writer, count, " ")

	//slide over the window
	for i := 0; i < (n - k); i++ {
		_, ok := vowel[str[i]]
		if ok {
			count--
		}

		_, ok = vowel[str[k+i]]
		if ok {
			count++
		}

		fmt.Fprint(writer, count, " ")
	}
}
