package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().Unix()))

func PrivateKey(p *big.Int) *big.Int {
	out := big.NewInt(1)
	// if p == 1
	if p.Cmp(big.NewInt(1)) == 0 {
		return out
	}
	// generate random (0, p)
	out.Rand(rnd, p.Abs(p))

	// if generated number <= 1
	if out.Cmp(big.NewInt(1)) <= 0 {
		out.Add(out, big.NewInt(2))
	}
	return out
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	return a, PublicKey(a, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
