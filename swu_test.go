package swu

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/minio/sha256-simd"
	"github.com/stretchr/testify/assert"
)

var (
	c = elliptic.P256()
)

func TestSWU(t *testing.T) {
	for i := 0; i < 10000; i++ {
		b := make([]byte, 32)
		rand.Read(b)

		x, y := HashToPoint(b)

		assert.True(t, elliptic.P256().IsOnCurve(x, y))
	}
}

func BenchmarkSWU(b *testing.B) {
	b.ResetTimer()
	buf := make([]byte, 32)
	rand.Read(buf)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		HashToPoint(buf)
	}
}

func BenchmarkTryInc(b *testing.B) {
	b.ResetTimer()
	buf := make([]byte, 32)
	rand.Read(buf)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		HashIntoCurvePoint(buf)
	}
}

func HashIntoCurvePoint(r []byte) (x, y *big.Int) {
	t := make([]byte, 32)
	copy(t, r)

	x, y = tryPoint(t)
	for y == nil || !c.IsOnCurve(x, y) {
		increment(t)
		x, y = tryPoint(t)

	}
	return
}

func tryPoint(r []byte) (x, y *big.Int) {
	hash := sha256.Sum256(r)
	x = new(big.Int).SetBytes(hash[:])

	// y² = x³ - 3x + b
	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)

	threeX := new(big.Int).Lsh(x, 1)
	threeX.Add(threeX, x)

	x3.Sub(x3, threeX)
	x3.Add(x3, c.Params().B)

	y = x3.ModSqrt(x3, c.Params().P)
	return
}

func increment(counter []byte) {
	for i := len(counter) - 1; i >= 0; i-- {
		counter[i]++
		if counter[i] != 0 {
			break
		}
	}
}
