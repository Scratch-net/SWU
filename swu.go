package swu

/*
 Implementation of Shallue-Woestijne-Ulas algorithm in Go
*/

import (
	"crypto/elliptic"
	"crypto/sha256"
	"math/big"
)

var (
	p   = elliptic.P256().Params().P
	a   *big.Int
	b   = elliptic.P256().Params().B
	mba *big.Int
)

func init() {
	f := &GF{p}
	a = f.Neg(Three)
	mba = f.Neg(f.Div(b, a))
}

func HashToPoint(data []byte) (x, y *big.Int) {
	f := &GF{p}

	hash := sha256.Sum256(data)

	t := new(big.Int).SetBytes(hash[:])
	t.Mod(t, p)

	//alpha = -t^2
	alpha := f.Neg(f.Square(t))

	// x2 = -(b / a) * (1 + 1/(alpha^2+alpha))
	x2 := f.Mul(mba, f.Add(One, f.Inv(f.Add(f.Square(alpha), alpha))))

	//x3 = alpha * x2
	x3 := f.Mul(alpha, x2)

	// h2 = x2^3 + a*x2 + b
	h2 := f.Add(f.Add(f.Cube(x2), f.Mul(a, x2)), b)
	// h3 = x3^3 + a*x3 + b
	h3 := f.Add(f.Add(f.Cube(x3), f.Mul(a, x3)), b)

	// tmp = h2 ^ ((p - 3) // 4)
	tmp := f.Pow(h2, f.Div(f.Sub(p, Three), Four))

	//if tmp^2 * h2 == 1:
	if f.Mul(f.Square(tmp), h2).Cmp(One) == 0 {
		// return (x2, tmp * h2 )
		return x2, f.Mul(tmp, h2)
	} else {
		//return (x3, h3 ^ ((p+1)//4))
		return x3, f.Pow(h3, f.Div(f.Add(p, One), Four))
	}
}
