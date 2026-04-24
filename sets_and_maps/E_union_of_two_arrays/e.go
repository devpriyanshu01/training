package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//n
	var n, m int
	fmt.Fscan(reader, &n)

	slice1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice1[i])
	}

	//m
	fmt.Fscan(reader, &m)

	slice2 := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &slice2[i])
	}

	//OPTIMISED SOLUTION
	union := make(map[int]struct{})
	
	for i := 0; i < n; i++ {
		union[slice1[i]] = struct{}{}
	}

	for i := 0; i < m; i++ {
		union[slice2[i]] = struct{}{}
	}

	length := len(union)

	ans := make([]int, length)
	i := 0
	for key :=  range union {
		ans[i] = key
		i++
	}

	sort.Ints(ans)
	
	fmt.Fprintln(writer, length)
	for _, val := range ans {
		fmt.Fprint(writer, val, " ")
	}

	//apply brute force - Gives TLE
	/*
	var union []int
	union = append(union, slice1[0])

	for i := 1; i < n; i++ {
		present := false
		for j := 0; j < len(union); j++ {
			if slice1[i] == union[j] {
				present = true
			}
		}
		if !present {
			union = append(union, slice1[i])
		}
	}

	for i := 0; i < m; i++ {
		present := false
		for j := 0; j < len(union); j++ {
			if slice2[i] == union[j] {
				present = true
			}
		}
		if !present {
			union = append(union, slice2[i])
		}
	}

	//sort the union
	sort.Ints(union)

	fmt.Fprintln(writer, len(union))
	for _, val := range union {
		fmt.Fprint(writer, val, " ")
	}
	*/
	
}
