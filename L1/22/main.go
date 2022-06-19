package main

import "fmt"

func calculate(a, b float64, action string) float64 {
	switch action {
	case "sum":
		return a + b
	case "division":
		return a / b
	case "multiply":
		return a * b
	case "subtract":
		return a - b
	default:
		println("Wrong action, choose for example 'sum', 'division', 'multiply', 'subsctract'")
		return 0
	}
}

func main() {
	fmt.Println(calculate(1048577.0, 1048578.0, "sum"))
	fmt.Println(calculate(1048577.0, 1048578.0, "division"))
	fmt.Println(calculate(1048577.0, 1048578.0, "multiply"))
	fmt.Println(calculate(1048577.0, 1048578.0, "subtract"))
}
