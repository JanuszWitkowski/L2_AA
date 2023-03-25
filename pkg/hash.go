package pkg

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"

	"golang.org/x/crypto/blake2b"
)

func bytesToUint(bytes []byte, hashLength uint) float64 {
	var sum float64 = 0.0
	fittingBytes := bytes[:hashLength]
	for _, b := range fittingBytes {
		sum += float64(b)
		sum /= float64(256)
	}

	return sum
}

func uintToBytes(data uint) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(data))

	return bytes
}

func Hash_sha1(data, hashLength uint) float64 {
	hash := sha1.New().Sum(uintToBytes(data))
	return bytesToUint(hash, hashLength)
}

func Hash_sha256(data, hashLength uint) float64 {
	hash := sha256.New().Sum(uintToBytes(data))
	return bytesToUint(hash, hashLength)
}

func Hash_blake2b(data, hashLength uint) float64 {
	hash := blake2b.Sum256(uintToBytes(data))
	return bytesToUint(hash[:], hashLength)
}
