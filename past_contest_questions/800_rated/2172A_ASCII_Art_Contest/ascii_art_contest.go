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

	aiScores := make([]int, 3)
	for i := range aiScores {
		fmt.Fscan(reader, &aiScores[i])
	}

	//sort the aiScores
	sort.Ints(aiScores)
	diff := aiScores[2] - aiScores[0]
	if diff >= 10 {
		fmt.Fprintln(writer, "check again")
	} else {
		fmt.Fprintln(writer, "final", aiScores[1])
	}
}
