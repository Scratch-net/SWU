package main

import "math/big"

type GF struct {
	P *big.Int
}

var (
	One   = big.NewInt(1)
	Two   = big.NewInt(2)
	Three = big.NewInt(3)
	Four  = big.NewInt(4)
)

func (g *GF) Neg(a *big.Int) *big.Int {
	return new(big.Int).Sub(g.P, a)
}

func (g *GF) Square(a *big.Int) *big.Int {
	return new(big.Int).Exp(a, Two, g.P)
}

func (g *GF) Cube(a *big.Int) *big.Int {
	return new(big.Int).Exp(a, Three, g.P)
}

func (g *GF) Pow(a, b *big.Int) *big.Int {
	return new(big.Int).Exp(a, b, g.P)
}

func (g *GF) Inv(a *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, g.P)
}

func (g *GF) Add(a, b *big.Int) *big.Int {
	add := new(big.Int).Add(a, b)
	return new(big.Int).Mod(add, g.P)
}

func (g *GF) Sub(a, b *big.Int) *big.Int {
	add := new(big.Int).Sub(a, b)
	return new(big.Int).Mod(add, g.P)
}

func (g *GF) Mul(a, b *big.Int) *big.Int {
	mul := new(big.Int).Mul(a, b)
	return new(big.Int).Mod(mul, g.P)
}

func (g *GF) Div(a, b *big.Int) *big.Int {
	invB := g.Inv(b)
	return g.Mul(a, invB)
}
