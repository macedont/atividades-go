package main

import (
	"fmt"
	"math"
)

/*
	interface é uma função que  extendida porém quem extender dela tem que fazer todos os metodos implementados nela.
*/

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func setArea(f forma) {
	fmt.Printf("A Área da forma é %.2f", f.area())
}

func main() {
	r := retangulo{10, 15}
	setArea(r)

	c := circulo{10}
	setArea(c)
}
