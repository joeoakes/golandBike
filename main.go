package main

import "fmt"

// Material enumeration
type Material int

const (
	STEEL Material = iota
	ALUMINUM
	CARBON_FIBER
)

// WheelType enumeration
type WheelType int

const (
	FRONT WheelType = iota
	REAR
)

// Wheel struct
type Wheel struct {
	Size float64
}

// Frame struct
type Frame struct {
	Material Material
}

// Bike struct 2
type Bike struct {
	Color  string
	Size   string
	Frame  Frame
	Wheels [2]Wheel
}

func (b *Bike) Start() {
	fmt.Println("Bike started!")
}

func (b *Bike) Stop() {
	fmt.Println("Bike stopped.")
}

func main() {
	myBike := Bike{
		Color: "Red",
		Size:  "Medium",
		Frame: Frame{Material: ALUMINUM},
		Wheels: [2]Wheel{
			{Size: 26.0},
			{Size: 26.0},
		},
	}

	fmt.Println("Bike color:", myBike.Color)
	fmt.Println("Bike size:", myBike.Size)
	fmt.Println("Frame material:", myBike.Frame.Material)
	fmt.Println("Wheel size:", myBike.Wheels[0].Size)

	myBike.Start()
	myBike.Stop()
}
