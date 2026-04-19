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
		var Ro, X, D, n int
		fmt.Fscan(reader, &Ro, &X, &D, &n)

		var contests string
		fmt.Fscan(reader, &contests)

		//logic begins here...
		ans := solution(contests, Ro, X, D, n)
		fmt.Fprintln(writer, ans)
	}
}

func solution(contests string, Ro, X, D, n int) int {
	var ans int
	for i := 0; i < n; i++ {
		//if contest is div1, it will rated irrespective of currRating
		if contests[i] == '1' {
			minPossible := Ro - D
			if minPossible > 0 {
				Ro -= D
				ans++
			}else {
				minPossibleNew := minPossible + (-1*minPossible) + 1
				Ro = minPossibleNew
				ans++
			}
		}else {
			//if contest is div 2
			if Ro < X {
				minPossible := Ro - D
				if minPossible > 0 {
					Ro -= D
					ans++
				}else {
					minPossibleNew := minPossible + (minPossible * -1) + 1
					Ro = minPossibleNew
					ans++
				}
			}
		}
	}
	return ans
}

