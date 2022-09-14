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
		var input string
		fmt.Fscan(in, &input)

		var current int
		var answer string

		for current < len(input) {
			if input[current] == '0' {
				answer += "a"
				current += 2
			} else {
				if input[current+1] == '1' {
					answer += "d"
					current += 2
				} else {
					if input[current+2] == '1' {
						answer += "c"
						current += 3
					} else {
						answer += "b"
						current += 3
					}
				}
			}
		}

		fmt.Fprintln(out, answer)
	}
}
