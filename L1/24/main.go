package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Point struct {
	x, y float64
}

type Constructor struct{}

func (r *Constructor) Construct() *Point {
	return &Point{x: rand.Float64() * 100.0, y: rand.Float64() * 100.0}
}

func FindDistance(p1, p2 *Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}

func main() {
	r := &Constructor{}
	p1 := r.Construct()
	p2 := r.Construct()
	fmt.Println(FindDistance(p1, p2))
}
