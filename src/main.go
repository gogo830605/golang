package main

import (
	// "routes"
	"fmt"
	"math"
	// "time"
)

	type geometry interface {
		area() float64
		perimeter() float64
	}
	
	type square struct {
	width, height float64
	}
	type circle struct {
	radius float64
	}
	
	func (s square) area() float64 {
	return s.width * s.height
	}
	func (s square) perimeter() float64 {
	return 2*s.width + 2*s.height
	}
	
	func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
	}
	func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
	}
	
	func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
	}

func main() {
	// router.SetupRouter()
	// router := router.SetupRouter()
	// router.Run(":8888")
	s := square{width: 3, height: 4}
// c := circle{radius: 5}

measure(s)
measure(s)
// measure(c)
}