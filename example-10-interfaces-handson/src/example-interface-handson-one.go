package main

import "fmt"

const pi = 3.14

//An interface is a blue print for every implementations. It has the collection of method definition
//Interface makes code flexible and scalable .
//In general application behaviour is defined in methods and those methods are defined in interfaces
type shape interface {
	getArea() float64
	getPerimeter() float64
}

type rectangle struct {
	hight, widht float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (rectang rectangle) getArea() float64 {
	return rectang.hight * rectang.widht
}

func (rectang rectangle) getPerimeter() float64 {
	return 2 * (rectang.hight + rectang.widht)
}

func (sq square) getArea() float64 {
	return sq.side * sq.side
}

func (sq square) getPerimeter() float64 {
	return 4 * (sq.side)
}

func (cr circle) getArea() float64 {
	return pi * cr.radius * cr.radius
}

func (cr circle) getPerimeter() float64 {
	return 2 * pi * (cr.radius)
}

func getDimensionsOfShape(sh shape) (float64, float64) {
	return sh.getArea(), sh.getPerimeter()
}

func main() {
	rect := rectangle{hight: 10, widht: 20}
	sqr := square{side: 10}
	cir := circle{radius: 7}

	rectarea, rectPerimeter := getDimensionsOfShape(rect)
	fmt.Println("Area of Rectangle --> ", rectarea)
	fmt.Println("Perimeter of Rectangle --> ", rectPerimeter)

	sqarea, sqPerimeter := getDimensionsOfShape(sqr)
	fmt.Println("Area of Square --> ", sqarea)
	fmt.Println("Perimeter of Square --> ", sqPerimeter)

	crarea, crPerimeter := getDimensionsOfShape(cir)
	fmt.Println("Area of Circle --> ", crarea)
	fmt.Println("Perimeter of Circle --> ", crPerimeter)

}
