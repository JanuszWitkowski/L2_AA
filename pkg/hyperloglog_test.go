package pkg

import (
	"fmt"
	"testing"
)

func TestTakeBBytes(t *testing.T) {
	tes := []byte{1, 34, 255, 255}
	hash := byteSliceToUint32(tes)

	n := takeBBites(11, hash)
	fmt.Print(n)
}

func TestTekaRestOfBytes(t *testing.T) {
	tes := []byte{1, 34, 255, 255}
	hash := byteSliceToUint32(tes)

	n := takeBBites(11, hash)
	b := takeRestOfBytes(11, hash)
	fmt.Print(n, b)
}

func TestRho(t *testing.T) {
	tes := []byte{1, 34, 255, 255}
	hash := byteSliceToUint32(tes)

	n := takeBBites(11, hash)
	b := takeRestOfBytes(11, hash)
	r := rho(b, 11)
	fmt.Print(n, b, r)
}

func TestHyperLogLog(t *testing.T) {
	ms := MultiSet_newMultiSet(4, 2)
	res := Hyperloglog(ms, Hash_blake2b_PURE, 5)

	fmt.Println(res)
}
