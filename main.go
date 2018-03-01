package goarea

import (
	"math"
)

const Pi = 3.1416

//Circ é publica
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi

}

//Rect é publica
func Rect(base, altura float64) float64 {
	return base * altura
}

//privada
func _Triang(base, altura float64) float64 {
	return (base * altura) / 2
}
