package util

import (
	"log"
)

type Color = Vec3

func NewColor(e0 float64, e1 float64, e2 float64) *Color {
	return &Color{[3]float64{e0, e1, e2}}
}

// Alias w sumie nie daje nic cieakwego, bo funkcja i tak jest na obiekcie
// to raczej powinno być utility function albo coś?
// Może nawet osobna klasa
func WriteColorTo(color *Color, logger *log.Logger) {
	r := color.X()
	g := color.Y()
	b := color.Z()

	rByte := int(255.999 * r)
	gByte := int(255.999 * g)
	bByte := int(255.999 * b)

	logger.Printf("%d %d %d\n", rByte, gByte, bByte)
}
