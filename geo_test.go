package main

import (
	"testing"
	"fmt"
)

func TestCoordNajac(t *testing.T) {
	coord, err := getCoordFromCity("najac")
	if err != nil {
		t.Fatalf("impossible d'extraire les coordonné de Najac: %v", err)
	}
	coodAttendue := Coordinates{Lat:44.22019, Lon:1.9809309}
	if coord != coodAttendue {
		t.Fatalf("Les coordonné de Najac ne sont pas correcte: %v", coord)
	}
	fmt.Printf("Coordonnées de Najac: %v", coodAttendue)
}