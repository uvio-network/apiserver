package sample

import (
	"crypto/rand"
	"errors"
	"io"
	"math/big"

	"github.com/uvio-network/apiserver/pkg/round"
	"github.com/xh3b4sd/tracer"
)

func (s *Sample) Random(a uint64, b uint64) ([]uint64, []uint64) {
	// At first we prepare the complete sequences of the lines we are asked to
	// select points from. Note that both lines may be of equal or different
	// length. In any case, having the complete sequence makes it easy to draw
	// from without allocating duplicates. Accounting to prevent duplicates after
	// the fact indicates bad algorithm design. So we prevent all of that
	// complexity by simply drawing from the complete sequences.

	var x []uint64
	for i := uint64(0); i < a; i++ {
		x = append(x, i)
	}

	var y []uint64
	for i := uint64(0); i < b; i++ {
		y = append(y, i)
	}

	// Randomize the complete sequences using any io.Reader that we got injected.
	// For testing purposes this may allow us to generate the same reliable
	// outcomes every time. In production we may use the global crypto/rand.Reader
	// as the source of cryptographically secure entropy.

	{
		mix(s.rea, x)
		mix(s.rea, y)
	}

	// Define the maximum number of points to return and map those points to an
	// output curve. We ensure a minimum of 1 point to return, and a maximum of 50
	// points to return each.

	var m uint64
	{
		m = min(a, b)
	}

	var n uint64
	{
		n = crv(m)
	}

	return lis(x, n), lis(y, n)
}

// crv returns point y on some curve given the minimum length m.
func crv(m uint64) uint64 {
	b := (m + 5) - ((m + 5) % 5)

	if m < 25 {
		return uint64(round.RoundP(float64(b)*0.20, 0))
	}

	if m < 100 {
		return 5
	}

	if m < 1000 {
		return uint64(round.RoundP(float64(m)*0.05, 0))
	}

	return 50
}

func lis(l []uint64, n uint64) []uint64 {
	if n > uint64(len(l)) {
		return nil
	}

	return l[:n]
}

func min(a uint64, b uint64) uint64 {
	if a > b {
		return b
	}

	return a
}

// mix is the crypto/rand version of math/rand.Shuffle. Below we use the given
// io.Reader, which is supposed to be some source of entropy, in order to
// shuffle the items of the given int slice.
func mix(rea io.Reader, lis []uint64) {
	if len(lis) == 0 {
		return
	}

	var m uint64
	{
		m = uint64(len(lis))
	}

	for i := m - 1; i > 0; i-- {
		b, err := rand.Int(rea, big.NewInt(int64(i+1)))
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			tracer.Panic(tracer.Mask(err))
		}

		{
			j := b.Int64()
			lis[i], lis[j] = lis[j], lis[i]
		}
	}
}
