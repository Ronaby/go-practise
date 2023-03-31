package main

import (
	"fmt"
	"math"
)

type Vertex7 struct {
	X, Y float64
}

func Abs(v Vertex7) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex7{3, 4}
	fmt.Println(Abs(v))
}
