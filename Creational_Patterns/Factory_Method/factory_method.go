package main

import "fmt"

type Shape interface {
	Draw()
}

type (
	ShapeFactory struct {
		shapeMap map[string]func() Shape
	}
	Circle struct{}

	Square struct{}

	Rectangle struct{}
)

func (c *Circle) Draw() {
	fmt.Println("Inside Circle::draw() method")
}

func (s *Square) Draw() {
	fmt.Println("Inside Square::draw() method")
}

func (r *Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method")
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		shapeMap: map[string]func() Shape{
			"circle":    func() Shape { return &Circle{} },
			"square":    func() Shape { return &Square{} },
			"rectangle": func() Shape { return &Rectangle{} },
		},
	}
}

func (f *ShapeFactory) CreateShape(shapeType string) Shape {
	if shapeFunc, ok := f.shapeMap[shapeType]; ok {
		return shapeFunc()
	}
	return nil
}

var (
	factory = NewShapeFactory()
	circle  = factory.CreateShape("circle")
	square  = factory.CreateShape("square")
	rect    = factory.CreateShape("rectangle")
)

func main() {
	circle.Draw()
	square.Draw()
	rect.Draw()
}
