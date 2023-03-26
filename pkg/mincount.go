package pkg

import (
	"sort"

	"golang.org/x/exp/slices"
)

func Mincount(ms *MultiSet, h func(uint, uint) float64, k uint, hashLength uint) float64 {
	M := make([]float64, k)

	for i := range M {
		M[i] = 1.0
	}

	for elem, err := ms.Next(); err == nil; elem, err = ms.Next() {
		helem := h(elem, hashLength)
		if float64(helem) < M[k-1] && !slices.Contains(M, helem) {
			M[k-1] = helem
			sort.Slice(M, func(i, j int) bool { return M[i] < M[j] })
		}
	}

	if M[k-1] == 1 {
		return float64(countNon1(M))
	}

	return float64(k-1) / M[k-1]
}

func countNon1(M []float64) uint {
	counter := 0

	for _, f := range M {
		if f != 1.0 {
			counter++
		}
	}

	return uint(counter)
}
