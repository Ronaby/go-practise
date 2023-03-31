package main

import (
	"fmt"
	"math"
)

type Vertex9 struct {
	X, Y float64
}

func Abs1(v Vertex9) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex9, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex9{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs1(v))
}
