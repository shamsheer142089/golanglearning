package main

import "fmt"

//An interface is a blue print for every implementations. It has the collection of method definition
//Interface makes code flexible and scalable .
//In general application behaviour is defined in methods and those methods are defined in interfaces
type shapes interface {
	getArea() float64
}

type triangle struct {
	hight, base float64
}

type squarePack struct {
	side float64
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.hight * tr.base
}

func (sq squarePack) getArea() float64 {
	return sq.side * sq.side
}

func getMeasurementOfShape(sh shapes) float64 {
	return sh.getArea()
}

func main() {
	trin := triangle{hight: 10, base: 20}
	sqr := squarePack{side: 10}

	trArea := getMeasurementOfShape(trin)
	fmt.Println("Area of Triangle --> ", trArea)

	sqarea := getMeasurementOfShape(sqr)
	fmt.Println("Area of Square --> ", sqarea)

}
