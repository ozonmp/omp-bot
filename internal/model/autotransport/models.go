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

func (g *Ground) ValidateFields() error {
	switch {
	case len(g.Name) == 0, len(g.Color) == 0:
		return fmt.Errorf("error: Name and Color cannot be empty")
	case g.MaxSpeed <= 0, g.WheelsCount <= 0:
		return fmt.Errorf("error: WheelsCount and MaxSpeed must be greater than zero")
	}
	return nil
}

func (g *Ground) Copy(ground Ground) bool {
	updatedFieldsCount := 0

	if len(ground.Name) > 0 {
		g.Name = ground.Name
		updatedFieldsCount++
	}

	if len(ground.Color) > 0 {
		g.Color = ground.Color
		updatedFieldsCount++
	}

	if ground.WheelsCount > 0 {
		g.WheelsCount = ground.WheelsCount
		updatedFieldsCount++
	}

	if ground.MaxSpeed > 0 {
		g.MaxSpeed = ground.MaxSpeed
		updatedFieldsCount++
	}

	return updatedFieldsCount > 0
}
