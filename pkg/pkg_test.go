package pkg

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestHashToFloat(t *testing.T) {
	testSlice := []byte{4, 7, 8, 9, 1, 5}

	fmt.Println(testSlice[:3])
	fmt.Println(testSlice[len(testSlice)-3:])
	fmt.Println("asfsa")
}

func TestHash(t *testing.T) {
	xd := sha1.Sum(uintToBytesLE(100))
	fmt.Println(uintToBytesLE(100))
	for _, x := range xd {
		fmt.Printf("%x", x)
	}

	fmt.Println()
}
