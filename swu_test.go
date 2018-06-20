package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSWU(t *testing.T) {
	for i := 0; i < 10000; i++ {
		b := make([]byte, 32)
		rand.Read(b)

		x, y := HashToPoint(b)

		assert.True(t, elliptic.P256().IsOnCurve(x, y))
	}
}
