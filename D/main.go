package main

import (
	"bufio"
	"fmt"
	"os"
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

		fmt.Fprintln(out, solve(input, inputArray))
	}
}

func solve(elements int, array []int) int {
	if elements == 1 || elements == 2 {
		return elements
	}

	var max, counter, one, two int

	for i := 0; i < elements; i++ {
		currentClient := array[i]

		if one == 0 || currentClient == one {
			one = currentClient
			counter++
			if i == len(array)-1 && max < counter {
				max = counter
			}
			continue
		} else if two == 0 {
			two = currentClient
			counter++
			if i == len(array)-1 && max < counter {
				max = counter
			}
			continue
		} else {
			if currentClient == one || currentClient == two {
				counter++
				if i != elements-1 {
					continue
				}
			}
		}

		if max < counter {
			max = counter
		}

		counter = 0
		one = 0
		two = 0
		i--

		for array[i] == array[i-1] {
			i--
		}
		i--
	}
	if max == 0 {
		return counter
	}

	return max
}
