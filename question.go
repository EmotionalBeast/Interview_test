package main

import "fmt"

func getDays(year, month uint) uint {
	var days uint
	isLeapYear := false
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		isLeapYear = true
	}
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if isLeapYear {
			days = 29
		} else {
			days = 28
		}
	default:
		days = 0
	}
	return days
}

func firstMonthCost(year, month, day, n uint) uint {
	monthDays := getDays(year, month)
	leftDays := monthDays - day + 1
	if leftDays > 30 {
		return n
	}
	return uint(float64(leftDays) / 30.0 * float64(n))
}

func main() {
	fmt.Println(firstMonthCost(2020, 8, 18, 60))
	fmt.Println(firstMonthCost(2020, 8, 1, 30))
}
