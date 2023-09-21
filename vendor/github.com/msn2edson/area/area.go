package area

import "math"

// Pi eh uma constante
const Pi = 3.1446

// Circ calcula a area de uma circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nao eh visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
