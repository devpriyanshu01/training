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

	//no. of testcases
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		//total no. of elements
		var n int
		fmt.Fscan(reader, &n)

		//slice input
		cards := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &cards[i])
		}

		//let's color the cards
		// var redFirst bool //first card will be marked red

		c := true
		coloredCards := make([]rune, n)
		for i := 0; i < n; i++ {
			if c {
				coloredCards[i] = 'r'
			} else {
				coloredCards[i] = 'b'
			}
			c = !c
		}

		//check if coloredCards pattern is accepted
		color1 := checkColor(coloredCards)

		//create a map, card number pointing to its color
		numColor := make(map[int]rune)
		for i := 0; i < n; i++ {
			numColor[cards[i]] = coloredCards[i]
		}

		//copy the cards num array to a new array for reversing
		sortedCards := make([]int, n)
		for i := 0; i < n; i++ {
			sortedCards[i] = cards[i]
		}
		sort.Ints(sortedCards)

		//reversed color
		sortedColor := make([]rune, n)
		for i := 0; i < n; i++ {
			color := numColor[sortedCards[i]]
			sortedColor[i] = color
		}

		//validate the color pattern
		color1Reversed := checkColor(sortedColor)

		//let's try the situation when blue is choosen for first card
		/*
		c = true
		for i := 0; i < n; i++ {
			if c {
				coloredCards[i] = 'b'
			} else {
				coloredCards[i] = 'r'
			}
			c = !c
		}
		//validate this pattern
		color2 := checkColor(coloredCards)

		//prepare the map of integer & color again
		//delete all the elements from the map
		for key := range numColor {
			delete(numColor, key)
		}

		//re-insert the value in the map
		for i := 0; i < n; i++ {
			numColor[cards[i]] = coloredCards[i]
		}

		//we already have sortedCards array
		//now get the new color pattern
		for i := 0; i < n; i++ {
			sortedColor[i] = numColor[sortedCards[i]]
		}

		//validate this pattern as well
		color2Reversed := checkColor(sortedColor)


		//answer decider
		if (color1 && color1Reversed) || (color2 && color2Reversed) {
			fmt.Fprintln(writer, "Yes")
			} else {
				fmt.Fprintln(writer, "No")
			}
			*/

		if color1 && color1Reversed {
			fmt.Fprintln(writer, "Yes")
		}else {
			fmt.Fprintln(writer, "No")
		}

	}
}

func checkColor(colorSlice []rune) bool {
	first := colorSlice[0]

	for i := 1; i < len(colorSlice); i++ {
		if colorSlice[i] == first {
			return false
		}
		first = colorSlice[i]
	}
	return true
}
