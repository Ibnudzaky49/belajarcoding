package main

func factorialLoop(value int) int {
	// result := 1
	// for i := value; 1 > 0; i-- {
	// result *= i
	// }
	// return result

	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}
