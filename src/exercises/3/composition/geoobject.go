package main

import "fmt"

type Position struct {
	x, y int
}

type GeoObject struct {
	pos   Position
	color int
}

type Painter interface {
	Paint() string
}

type Circle struct {
	GeoObject
	radius int
}

type Rectangle struct {
	GeoObject
	width, height int
}

type Triangle struct {
	GeoObject
	p1, p2, p3 Position
}

func (g GeoObject) Paint() (result string) {
	result = fmt.Sprintf("Name: %T, PosX : %d, PosY: %d, Color: %d", g, g.pos.x, g.pos.y, g.color)
	return
}

func (c Circle) Paint() (result string) {
	parentResult := c.GeoObject.Paint()
	result = parentResult + fmt.Sprintf(", Radius: %d", c.radius)
	return
}

func (r Rectangle) Paint() (result string) {
	parentResult := r.GeoObject.Paint()
	result = parentResult + fmt.Sprintf(", Width: %d, Height: %d", r.width, r.height)
	return
}

func (t Triangle) Paint() (result string) {
	parentResult := t.GeoObject.Paint()
	result = parentResult + fmt.Sprintf(", A: %d, B: %d, C: %d", t.p1, t.p2, t.p3)
	return
}

func main() {
	// Polymorph slice of Painter objects
	objects := []Painter{
		// short initialization
		Circle{GeoObject{Position{1, 2}, 3}, 40},
		Rectangle{GeoObject{Position{1, 2}, 4}, 10, 10},
		Triangle{GeoObject{Position{1, 2}, 3}, Position{10, 20}, Position{11, 21}, Position{12, 22}},

		// or with named identifiers
		Circle{
			GeoObject: GeoObject{
				pos:   Position{x: 1, y: 2},
				color: 3},
			radius: 40},
		Rectangle{
			GeoObject: GeoObject{
				pos:   Position{x: 1, y: 2},
				color: 4},
			width:  10,
			height: 10},
		Triangle{
			GeoObject: GeoObject{
				pos:   Position{x: 1, y: 2},
				color: 3},
			p1: Position{x: 10, y: 20},
			p2: Position{x: 11, y: 21},
			p3: Position{x: 12, y: 22}},
	}
	for _, v := range objects {
		fmt.Printf(v.Paint() + "\n")
	}
}
