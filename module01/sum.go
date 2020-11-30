package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	sum := 0
	for _, item := range numbers {
		sum += item
	}

	return sum
}

// Sum - recursive solution
// func Sum(numbers []int) int {
// 	if len(numbers) == 0 {
// 		return 0
// 	}

// 	return numbers[0] + Sum(numbers[1:])
// }

/*
Sum(3,4,5) =>
	3 + Sum(4,5) =>
		4 + Sum(5) =>
			5 + Sum() =>
				0
*/

// Run only this test:
// $ gotest -run=TestSum -v
