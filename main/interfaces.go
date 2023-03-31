package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func (f MyFloat) Abs1() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex10 struct {
	X, Y float64
}

func (v *Vertex10) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex10{3, 4}

	a = f
	a = &v

	a = &v
	fmt.Println(a.Abs())
}
