package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0

	for {
		fmt.Println("\n**************Main Menu**********************\n")
		fmt.Println("1 Calculate Area")
		fmt.Println("2 Calculate Volume")
		fmt.Println("0 Exit")
		if _, err := fmt.Scan(&x); err == nil {
			switch x {
			case 1:
				y := 0
				for {
					fmt.Println("\n**************Sub Menu******Calculate Area****\n")
					fmt.Println("1 Rectangle")
					fmt.Println("2 Circle")
					fmt.Println("9 Main Menu")
					fmt.Println("0 Exit")
					if _, err := fmt.Scan(&y); err == nil {
						switch y {
						case 1:
							areaRectangle()
						case 2:
							areaCircle()
						case 9:
							goto exitArea
						case 0:
							goto outerLoop
						default:
							fmt.Println("\nInvalid Choice. Enter again")
						}
					}
				}
			exitArea:
				break
			case 2:
				y := 0
				for {
					fmt.Println("\n**************Sub Menu******Calculate Volume****\n")
					fmt.Println("1 Cylinder")
					fmt.Println("2 Cube")
					fmt.Println("9 Main Menu")
					fmt.Println("0 Exit")
					if _, err := fmt.Scan(&y); err == nil {
						switch y {
						case 1:
							volumeCylinder()
						case 2:
							volumeCube()
						case 9:
							goto exitVolume
						case 0:
							goto outerLoop
						default:
							fmt.Println("\nInvalid Choice. Enter again")
						}
					}
				}
			exitVolume:
				break
			case 0:
				goto outerLoop
			default:
				fmt.Println("\nInvalid Choice. Enter again")
			}
		}
	}
outerLoop:
}

func areaCircle() {
	var r float32
	fmt.Println("\nEnter the radius of the circle:")
	if _, err := fmt.Scan(&r); err == nil {
		fmt.Printf("\nArea of the circle of radius%5.2f is: %5.2f\n", r, math.Pi*r*r)
	}
}

func areaRectangle() {
	var l, w float32
	fmt.Println("\nEnter the length and width of the rectangle: ")
	if _, err := fmt.Scan(&l, &w); err == nil {
		fmt.Printf("\nArea of the rectangle %5.2f x %5.2f is: %5.2f\n", l, w, l*w)
	}
}

func volumeCylinder() {
	var r, h float32
	fmt.Println("\nEnter the radius and height of the cylinder: ")
	if _, err := fmt.Scan(&r, &h); err == nil {
		fmt.Printf("\nVolume of the cylinder of radius %5.2f and height %5.2f is: %5.2f\n", r, h, math.Pi*r*r*h)
	}
}

func volumeCube() {
	var s float32
	fmt.Println("\nEnter the side of the cube: ")
	if _, err := fmt.Scan(&s); err == nil {
		fmt.Printf("\nVolume of the cube of side %5.2f is: %5.2f\n", s, s*s*s)
	}
}
