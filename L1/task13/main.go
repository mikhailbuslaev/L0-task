package main

func main() {
	a, b := 0, 1
	println(a, b)
	a, b = b, a
	println(a, b)
}
