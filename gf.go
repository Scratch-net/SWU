package swu

import (
	"math/big"
	"sync"
)

type GF struct {
	P *big.Int
}

var (
	One   = big.NewInt(1)
	Two   = big.NewInt(2)
	Three = big.NewInt(3)
	Four  = big.NewInt(4)
	pool  = sync.Pool{New: func() interface{} {
		return new(big.Int)
	}}
)

func (g *GF) NewInt() *big.Int {
	return pool.Get().(*big.Int)
}

func (g *GF) FreeInt(i ...*big.Int) {
	for _, x := range i {
		pool.Put(x)
	}

}

func (g *GF) Neg(a *big.Int) *big.Int {
	return g.NewInt().Sub(g.P, a)
}

func (g *GF) Square(a *big.Int) *big.Int {
	return g.NewInt().Exp(a, Two, g.P)
}

func (g *GF) Cube(a *big.Int) *big.Int {
	return g.NewInt().Exp(a, Three, g.P)
}

func (g *GF) Pow(a, b *big.Int) *big.Int {
	return g.NewInt().Exp(a, b, g.P)
}

func (g *GF) Inv(a *big.Int) *big.Int {
	return g.NewInt().ModInverse(a, g.P)
}

func (g *GF) Add(a, b *big.Int) *big.Int {
	add := g.NewInt().Add(a, b)
	return add.Mod(add, g.P)
}

func (g *GF) Sub(a, b *big.Int) *big.Int {

	negB := g.NewInt().Sub(a, b)
	return negB.Mod(negB, g.P)
}

func (g *GF) Mul(a, b *big.Int) *big.Int {
	mul := new(big.Int).Mul(a, b)
	return mul.Mod(mul, g.P)
}

func (g *GF) Div(a, b *big.Int) *big.Int {
	invB := g.Inv(b)
	t := g.Mul(a, invB)
	g.FreeInt(invB)
	return t
}
