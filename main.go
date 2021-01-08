package main

import "fmt"

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

	return fmt.Sprintf("%s%.0fkg %.2fm/s%s", prefix, b.m, b.v, suffix)
}

func main() {
	b1 := &Body{1, 0}
	b2 := &Body{100, -1}

	fmt.Println(b1, b2)

	var cnt int
	for true {
		if b2.v < 0 || b1.v > b2.v {
			b2.Collision(b1)
			cnt++
			fmt.Printf("%d: %s %s\n", cnt, b1, b2)
		} else {
			break
		}

		if b2.v > 0 {
			fmt.Println("-------------------------------------------")
		}

		if b1.v < 0 {
			b1.WallCollision()
			cnt++
			fmt.Printf("%dw: %s %s\n", cnt, b1, b2)
		}
	}

	fmt.Printf("Pi=%f\n", float64(cnt)/1.0)
}
