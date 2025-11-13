package const_und_enums

import (
	"fmt"
)

// 0: Norden
// 1: Osten
// 2: Süden
// 3: Westen
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

/*const (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)*/

type Coord struct {
	Longitude float64
	Latitude  float64
}

func (c *Coord) Step(d Direction) {
	switch d {
	case North:
		c.Latitude += 1
	case East:
		c.Longitude += 1
	case South:
		c.Latitude -= 1
	case West:
		c.Longitude -= 1
	}

	/*
		if d == 0 {
			c.Latitude += 1
		} else if d == 1 {
			c.Longitude += 1
		} else if d == 2 {
			c.Latitude -= 1
		} else if d == 3 {
			c.Longitude -= 1
		}
	*/
}

func ExampleCoord_usage() {
	c1 := Coord{1.5, 3}
	c2 := Coord{Latitude: 4, Longitude: 5.5}

	fmt.Println(c1)
	fmt.Println(c2)

	c1.Step(0)
	c1.Step(1)

	fmt.Println(c1)

	c1.Step(North)
	c1.Step(East)

	fmt.Println(c1)

	// Output:
	// {1.5 3}
	// {5.5 4}
	// {2.5 4}
	// {3.5 5}
}

/* besseres Step
func (c *Coord) Step(d Direction) error {
	switch d {
	case 0:
		c.Latitude += 1
	case 1:
		c.Longitude += 1
	case 2:
		c.Latitude -= 1
	case 3:
		c.Longitude -= 1
	default:
		return fmt.Errorf("Geht nicht, du ...")
	}
	return nil
*/
