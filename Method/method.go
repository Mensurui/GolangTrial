package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Sqrt(p.X*q.X + p.Y*q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p1 := Point{1, 3}
	p2 := Point{3, 4}
	var rf = 3.00

	fmt.Println("The distance from p1 to p2 is: ", p1.Distance(p2))
	p1.ScaleBy(rf)
	fmt.Println("\nRefactored distance: ", p1)

	red := color.RGBA{225, 0, 0, 225}
	blue := color.RGBA{0, 0, 225, 225}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))

	p.ScaleBy(2)
	q.ScaleBy(3)
	fmt.Println(p.Distance(q.Point))
}
