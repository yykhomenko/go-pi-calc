package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type body struct {
	m float64 // [kg]
	v float64 // [m/s]
}

func (b body) p() float64 {
	return b.m * b.v // [kg*m/s]
}

func (b *body) collision(b2 *body) {
	v1 := (2*b2.v*b2.m + b.v*(b.m-b2.m)) / (b.m + b2.m)
	b2.v = (2*b.v*b.m + b2.v*(b2.m-b.m)) / (b.m + b2.m)
	b.v = v1
}

func (b *body) wallCollision() {
	b.v = -b.v
}

func calcPi(d int) float64 {
	b1 := &body{1, 0}
	b2 := &body{math.Pow(100, float64(d-1)), -1}

	var cnt int
	for b2.v < 0 || b1.v > b2.v {
		b2.collision(b1)
		cnt++
		if b1.v < 0 {
			b1.wallCollision()
			cnt++
		}
	}

	return float64(cnt) / math.Pow(10, float64(d-1))
}

func main() {
	digits, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(calcPi(digits))
}
