package ground

import "fmt"

var allGrounds = []Ground{
	{Name: "auto", WheelsCount: 4, Color: "black", MaxSpeed: 200},
	{Name: "bus", WheelsCount: 4, Color: "blue", MaxSpeed: 120},
	{Name: "bike", WheelsCount: 2, Color: "red", MaxSpeed: 15},
	{Name: "motorbike", WheelsCount: 2, Color: "black", MaxSpeed: 80},
	{Name: "scooter", WheelsCount: 2, Color: "yellow", MaxSpeed: 25},
}

type Ground struct {
	Name        string
	WheelsCount int
	Color       string
	MaxSpeed    int
}

func (c Ground) String() string {
	return fmt.Sprintf("%s, \tWheelsCount: %d, \tColor: %s, \tMaxSpeed: %d",
		c.Name, c.WheelsCount, c.Color, c.MaxSpeed)
}

func NewGround(name string, wheelsCount int, color string, maxSpeed int) Ground {
	return Ground{
		Name:        name,
		WheelsCount: wheelsCount,
		Color:       color,
		MaxSpeed:    maxSpeed,
	}
}
