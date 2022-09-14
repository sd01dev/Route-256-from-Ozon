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

		var rows, columns int
		fmt.Fscan(in, &rows, &columns)

		visited := make([][]int, rows)

		for i := 0; i < rows; i++ {
			var tmp string
			visited[i] = make([]int, columns)
			fmt.Fscan(in, &tmp)

			for j := range tmp {
				if tmp[j] == '.' {
					visited[i][j] = 0
				} else {
					visited[i][j] = 1
				}
			}
		}

		var start bool
		var answer string
		var neighbors int
		var currentRow, currentColumn int

		// Поиск начала
		for i := 0; i < len(visited); i++ {
			if start == true {
				break
			}

			for j := 0; j < len(visited[i]); j++ {
				neighbors = 0
				if visited[i][j] == 1 {
					if i+1 < len(visited) && visited[i+1][j] == 1 {
						neighbors++
					}
					if i != 0 && visited[i-1][j] == 1 {
						neighbors++
					}
					if j+1 < len(visited[i]) && visited[i][j+1] == 1 {
						neighbors++
					}
					if j != 0 && visited[i][j-1] == 1 {
						neighbors++
					}
				}

				if neighbors == 1 {
					currentRow = i
					currentColumn = j
					start = true
					break
				}
			}
		}

		previousRow, previousColumn := 100000, 1000000

		for {
			if start {
				if currentRow+1 != previousRow && currentRow+1 < len(visited) && visited[currentRow+1][currentColumn] == 1 {
					answer += "D"
					previousRow, previousColumn = currentRow, currentColumn
					currentRow++
					continue
				}
				if currentRow-1 != previousRow && currentRow != 0 && visited[currentRow-1][currentColumn] == 1 {
					answer += "U"
					previousRow, previousColumn = currentRow, currentColumn
					currentRow--
					continue
				}
				if currentColumn+1 != previousColumn && currentColumn+1 < len(visited[currentRow]) && visited[currentRow][currentColumn+1] == 1 {
					answer += "R"
					previousRow, previousColumn = currentRow, currentColumn
					currentColumn++
					continue
				}
				if currentColumn-1 != previousColumn && currentColumn != 0 && visited[currentRow][currentColumn-1] == 1 {
					answer += "L"
					previousRow, previousColumn = currentRow, currentColumn
					currentColumn--
					continue
				}

				start = false
			} else {
				break
			}
		}

		fmt.Fprintln(out, answer)
	}
}
