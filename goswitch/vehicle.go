package main

import "fmt"

type vehicle interface {
	printVehicleDetails(string)
}

type BaseVehicleData struct {
	model string
	make  string
	color string
	cc    float32
	year  int
}

type ExtendedVehicleData struct {
	BaseVehicleData
	typeVehicle string
	seating     int
	doors       int
}

type car struct {
	ExtendedVehicleData
}

type truck struct {
	ExtendedVehicleData
	loadingCapacity float32
}

type bike struct {
	BaseVehicleData
}

var inventory []vehicle

func init() {
	var timeMachine vehicle

	inventory = []vehicle{
		bike{BaseVehicleData{"FTR 1200", "Indian", "Black", 1203.0, 2019}},
		bike{BaseVehicleData{"Iron 1200", "Harley", "Blue", 1200.0, 2018}},
		car{ExtendedVehicleData{BaseVehicleData{"Sonata", "Hyundai", "White", 2400.0, 2017}, "Sedan", 5, 4}},
		car{ExtendedVehicleData{BaseVehicleData{"Santa Fe", "Hyundai", "Red", 2400.0, 2016}, "SUV", 7, 4}},
		car{ExtendedVehicleData{BaseVehicleData{"Civic", "Honda", "White", 2000.0, 2017}, "Hatchback", 5, 4}},
		car{ExtendedVehicleData{BaseVehicleData{"A5", "Audi", "Red", 3000.0, 2019}, "Coupe", 2, 2}},
		car{ExtendedVehicleData{BaseVehicleData{"Mazda 6", "Mazda", "White", 2500.0, 2018}, "Sedan", 5, 4}},
		car{ExtendedVehicleData{BaseVehicleData{"CRV", "Honda", "Black", 1500.0, 2017}, "SUV", 5, 4}},
		car{ExtendedVehicleData{BaseVehicleData{"Camry", "Toyota", "Silver", 3500.0, 2018}, "Sedan", 5, 4}},
		truck{ExtendedVehicleData{BaseVehicleData{"F-150", "Ford", "Gray", 3600.0, 2014}, "Truck", 7, 4}, 13200.0},
		truck{ExtendedVehicleData{BaseVehicleData{"RAM1500", "Dodge", "White", 5700.0, 2017}, "Truck", 5, 2}, 12750.0},
		truck{ExtendedVehicleData{BaseVehicleData{"Silverado", "Chevrolet", "Black", 6000.0, 2016}, "Truck", 7, 4}, 14500.0},
		timeMachine,
	}
}

func main() {
	carCount, truckCount, bikeCount := 0, 0, 0
	fmt.Printf("%-15v%-12v%-12v%-10v%-4v%12v%12v%8v%12v\n", "Type", "Make", "Model", "Color", "CC", "Year", "Seating", "Doors", "Load(lbs)")
	fmt.Println("----------------------------------------------------------------------------------------------")

	for _, veh := range inventory {
		switch v := veh.(type) {
		case car:
			v.carDetails()
			carCount++
		case truck:
			v.truckDetails()
			truckCount++
		case bike:
			v.bikeDetails()
			bikeCount++
		default:
			fmt.Printf("Are you sure this vehicle type exists")
		}
	}
}

func (c *car) carDetails() {
	msg := fmt.Sprintf("Car-%-11v%-12v%-12v%-10v%-6.2f%9v%9d%9d\t--\n", c.typeVehicle, c.make, c.model, c.color, c.cc, c.year, c.seating, c.doors)
	c.printVehicleDetails(msg)
}

func (t *truck) truckDetails() {
	msg := fmt.Sprintf("Truck-%-9v%-12v%-12v%-10v%-6.2f%9v%9d%9d%12.2f\n", t.typeVehicle, t.make, t.model, t.color, t.cc, t.year, t.seating, t.doors, t.loadingCapacity)
	t.printVehicleDetails(msg)
}

func (b *bike) bikeDetails() {
	msg := fmt.Sprintf("Bike           %-12v%-12v%-10v%-6.2f%9v\t--\t  --\t--\n", b.make, b.model, b.color, b.cc, b.year)
	b.printVehicleDetails(msg)
}

func (c car) printVehicleDetails(s string) {
	fmt.Printf("%v", s)
}

func (t truck) printVehicleDetails(s string) {
	fmt.Printf("%v", s)
}

func (b bike) printVehicleDetails(s string) {
	fmt.Printf("%v", s)
}
