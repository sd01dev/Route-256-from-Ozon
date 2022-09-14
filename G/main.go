package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var num int
	fmt.Fscan(in, &num)

	for k := 1; k <= num; k++ {
		var input int
		fmt.Fscan(in, &input)
		inputArray := make([]int, input)
		for i := 0; i < input; i++ {
			var value int
			fmt.Fscan(in, &value)
			inputArray[i] = value
		}

		var sortedArray []int
		sortedArray = append(sortedArray, inputArray...)

		sort.Ints(sortedArray)

		reqNum := 1
		var currentNum int
		stars := 5
		starMap := make(map[int]int)

		for i := len(sortedArray) - 1; i >= 0; i-- {
			if currentNum >= reqNum && sortedArray[i] != sortedArray[i+1] {
				if stars != 1 {
					stars--
					reqNum = currentNum + 1
					currentNum = 0
				}
			}

			_, ok := starMap[sortedArray[i]]
			if !ok {
				starMap[sortedArray[i]] = stars
				currentNum++
			} else {
				currentNum++
				continue
			}

		}

		if stars != 1 || currentNum < reqNum {
			for i, _ := range inputArray {
				inputArray[i] = -1
			}
		} else {
			for i, value := range inputArray {
				inputArray[i] = starMap[value]
			}
		}

		for _, value := range inputArray {
			fmt.Fprintf(out, "%d ", value)
		}
		fmt.Fprintln(out, "")
	}
}
