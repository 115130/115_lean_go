package main

type Rectangle struct {
	weight float64
	height float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return (rectangle.weight + rectangle.height) * 2
}

func Area(rectangle Rectangle) (perimeter float64) {
	return rectangle.weight * rectangle.height
}
