/* Program that interacts with shapes */

/* Calculating the area of basic shapes e.g Circle, Rectangle */

package main

import ("fmt"; "math")

/* Struct for Circle */
type Circle struct {
	x, y, r float64
}

/* Struct for Rectangle */
type Rectangle struct {
	x1, y1, x2, y2 float64
}

/* Calculation of distance */
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

/* Area of the rectangle */
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

/* Area of the circle */
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

/* This is the main function */
func main() {
	c := Circle{0, 0, 5}

	r := Rectangle{0, 0, 10, 10}

	fmt.Println("Area of the rectangle is: ", r.area())

	fmt.Println("Aread of the circle is: ", c.area())
}