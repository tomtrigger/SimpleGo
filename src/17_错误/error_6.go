package main

import (
	"fmt"
	"errors"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		// Area calculation failed, radius is less than zero
		return 
	}
	fmt.Printf("Area of circle %0.2f", area)
}