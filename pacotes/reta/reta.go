package main

import "math"

// Iniciando com letra maiuscula eh PUBLICO (visibilidade dentro e fora do pacote)
// Iniciando com letra minuscula eh PRIVADO (visibilidade somente dentro do pacote)

// Por exemplo...
// Ponto - gerara algo publico
// ponto ou _Ponto - gerara algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return cx, cy
}

// Distancia eh responsavel por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
