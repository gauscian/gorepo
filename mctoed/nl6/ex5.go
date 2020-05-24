/*
* Created on Mon May 18 2020
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
	circle area= π r 2
	square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA
	method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
* Primary Author - Achin Gupta
* Copyright (c) 2020 Netskope
*/

package main

import "fmt"

// Shape is any shape
type Shape interface {
	area() float64
}

// Square is a square
type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

// Circle is a circle
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return (22/7) * c.radius * c.radius
}

// Area of any shape
func Area(s Shape) {
	fmt.Println(s.area())
}

func main() {
	// polymorphism
	Area(Square{3})
	Area(Circle{2.3})
}
