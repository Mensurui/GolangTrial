package main

import (
	"fmt"
)

type Points struct {
	X float64
	Y float64
}

func (p Points) geo(q Points) {
	fmt.Println(p.X, q.X)
}

func ma() {
	p := Points{1, 2}
	q := Points{2, 3}
	fmt.Println(p.geo(q))
}
