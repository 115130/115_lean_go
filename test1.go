package main

import "fmt"

type P interface {
	Test() float64
}
type type1 struct {
	height float64
	weight float64
}

func (in type1) Test() float64 {
	return in.height * in.weight
}
func main() {
	p := type1{10, 10}
	a := p.Test()
	fmt.Printf("%.10f", a)
}
