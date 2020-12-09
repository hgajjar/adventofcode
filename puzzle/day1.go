package puzzle

// ReportRepairPart1 finds the two numbers from input which adds up to 2020
// and returns their product
func ReportRepairPart1(expenseReport []int) int {
	for x, a := range expenseReport {
		for _, b := range expenseReport[x+1:] {
			if a+b == 2020 {
				return a * b
			}
		}
	}
	panic("Can not find the result")
}

// ReportRepairPart2 finds the three numbers from input which adds up to 2020
// and returns their product
func ReportRepairPart2(expenseReport []int) int {
	for x, a := range expenseReport {
		for y, b := range expenseReport[x+1:] {
			for _, c := range expenseReport[y+1:] {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	panic("Can not find the result")
}
