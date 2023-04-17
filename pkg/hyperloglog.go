package pkg

import (
	"math"
	"math/bits"
)

func Hyperloglog(ms *MultiSet, h func(uint, uint) []byte, b uint32) float64 {
	m := uint(math.Pow(2, float64(b)))

	M := make([]uint32, m)

	for i := range M {
		M[i] = 0
	}

	for elem, err := ms.Next(); err == nil; elem, err = ms.Next() {
		helem := h(elem, 4)
		hash := byteSliceToUint32(helem)
		j := takeBBites(b, hash)
		w := takeRestOfBytes(b, hash)
		M[j] = max(M[j], rho(w, b))

	}
	sumWeird := 0.0

	for _, elem := range M {
		sumWeird += math.Pow(2.0, -float64(elem))
	}

	nHLL := getAlpha(ms.n*ms.m) * float64(m*m) * (1.0 / sumWeird)

	if nHLL < 2.5*float64(m) {
		V := count0(M)
		if V != 0 {
			nHLL = float64(m) * math.Log(float64(m)/float64(V))
		}
	}

	TTTT := math.Pow(2.0, 32.0)
	if nHLL > (1.0/30.0)*TTTT {
		nHLL = -TTTT * math.Log(1-(nHLL/TTTT))
	}
	return nHLL
}

func takeBBites(b uint32, hash uint32) uint32 {
	return hash >> (32 - b)
}

func takeRestOfBytes(b uint32, hash uint32) uint32 {
	return hash << b
}

func byteSliceToUint32(bytes []byte) uint32 {
	var sum uint32 = 0

	for _, b := range bytes {
		sum = sum << 8
		sum += uint32(b)
	}

	return sum
}

func rho(hash uint32, bit uint32) uint32 {
	if hash == 0 {
		return 32 - bit + 1
	}
	return uint32(bits.LeadingZeros32(hash)) + 1
}

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}

	return b
}

func getAlpha(m uint) float64 {
	if m < 128 {
		return 0.7
	}

	return 0.7213 / (1.0 + 1.079/float64(m))
}

func count0(M []uint32) uint {
	counter := 0

	for _, f := range M {
		if f == 0 {
			counter++
		}
	}

	return uint(counter)
}
