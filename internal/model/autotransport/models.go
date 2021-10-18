package autotransport

import (
	"fmt"
	"strings"
)

var allGrounds = []Ground{
	{Name: "auto", WheelsCount: 4, Color: "black", MaxSpeed: 200},
	{Name: "bus", WheelsCount: 4, Color: "blue", MaxSpeed: 120},
	{Name: "bike", WheelsCount: 2, Color: "red", MaxSpeed: 15},
	{Name: "motorbike", WheelsCount: 2, Color: "black", MaxSpeed: 80},
	{Name: "scooter", WheelsCount: 2, Color: "yellow", MaxSpeed: 25},
}

type Ground struct {
	Name        string
	WheelsCount uint64
	Color       string
	MaxSpeed    uint64
}

func (g Ground) String() string {
	return fmt.Sprintf("%s, \tWheelsCount: %d, \tColor: %s, \tMaxSpeed: %d",
		strings.Title(g.Name), g.WheelsCount, g.Color, g.MaxSpeed)
}

func AllGrounds() *[]Ground {
	return &allGrounds
}

func NewGround(name string, wheelsCount uint64, color string, maxSpeed uint64) Ground {
	return Ground{
		Name:        name,
		WheelsCount: wheelsCount,
		Color:       color,
		MaxSpeed:    maxSpeed,
	}
}
