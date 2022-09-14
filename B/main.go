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

		var day, month, year int
		fmt.Fscan(in, &day, &month, &year)

		if ifDateValid(ifLeapYear(year), day, month) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func ifDateValid(leapYear bool, day, month int) bool {
	if leapYear && month == 2 {
		return day < 30
	}

	if !leapYear && month == 2 {
		return day < 29
	}

	if month == 4 || month == 6 || month == 9 || month == 11 {
		return day < 31
	}

	return true
}

func ifLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else {
		return year%4 == 0 && year%100 != 0
	}
}
