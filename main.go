package main

import (
	"fmt"
	"math"
)

type Body struct {
	m float64 // [kg]
	v float64 // [m/s]
}

func (b Body) P() float64 {
	return b.m * b.v // [kg*m/s]
}

func (b *Body) Collision(b2 *Body) {
	v1 := (2*b2.v*b2.m + b.v*(b.m-b2.m)) / (b.m + b2.m)
	v2 := (2*b.v*b.m + b2.v*(b2.m-b.m)) / (b.m + b2.m)
	b.v = v1
	b2.v = v2
}

func (b *Body) WallCollision() {
	b.v = -b.v
}

func (b Body) String() string {
	var prefix, suffix string
	p := b.P()
	switch {
	case p < 0:
		prefix = "<"
		suffix = "|"
	case p > 0:
		prefix = "|"
		suffix = ">"
	default:
		prefix = "|"
		suffix = "|"
	}

	return fmt.Sprintf("%s%.0fkg %fm/s%s", prefix, b.m, b.v, suffix)
}

func calcPi(d int, debug bool) float64 {
	b1 := &Body{1, 0}
	b2 := &Body{math.Pow(100, float64(d-1)), -1}

	if debug {
		fmt.Println(b1, b2)
	}

	var cnt int
	var maxSpeed float64
	for true {

		if maxSpeed < b1.v {
			maxSpeed = b1.v
		}

		if b2.v < 0 || b1.v > b2.v {
			b2.Collision(b1)
			cnt++
			if debug {
				fmt.Printf("%d: %s %s\n", cnt, b1, b2)
			}
		} else {
			if debug {
				fmt.Printf("max speed of b1 = %fm/s\n", maxSpeed)
			}
			break
		}

		if b1.v < 0 {
			b1.WallCollision()
			cnt++
			if debug {
				fmt.Printf("%dw: %s %s\n", cnt, b1, b2)
			}
		}
	}

	return float64(cnt) / math.Pow(10, float64(d-1))
}

func main() {
	fmt.Printf("Pi=%f", calcPi(3, true))
}
