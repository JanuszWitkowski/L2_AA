package pkg

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
)

func bytesToFloat64LE(bytes []byte, hashLength uint) float64 {
	var sum float64 = 0.0
	fittingBytes := bytes[:hashLength]

	for i := len(fittingBytes) - 1; i >= 0; i-- {
		sum += float64(fittingBytes[i])
		sum /= float64(256)
	}

	return sum
}

// func bytesToFloat64BE(bytes []byte, hashLength uint) float64 {
// 	var sum float64 = 0.0
// 	meaningfullBytesStart := uint(len(bytes)) - hashLength
// 	fittingBytes := bytes[meaningfullBytesStart:]

// 	for i := 0; i < len(fittingBytes); i++ {
// 		sum += float64(fittingBytes[i])
// 		sum /= float64(256)
// 	}

// 	return sum
// }

func uintToBytesLE(data uint) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(data))

	return bytes
}

func uintToBytesBE(data uint) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(data))

	return bytes
}

func Hash_sha1(data, hashLength uint) float64 {
	hash := sha1.Sum(uintToBytesLE(data))
	return bytesToFloat64LE(hash[:], hashLength)
}

func Hash_sha256(data, hashLength uint) float64 {
	hash := sha256.Sum256(uintToBytesLE(data))
	return bytesToFloat64LE(hash[:], hashLength)
}

func Hash_sha3(data, hashLength uint) float64 {
	hash := sha3.Sum256(uintToBytesLE(data))
	return bytesToFloat64LE(hash[:], hashLength)
}

func Hash_blake2b(data, hashLength uint) float64 {
	hash := blake2b.Sum256(uintToBytesBE(data))
	return bytesToFloat64LE(hash[:], hashLength)
}

func Hash_blake2s(data, hashLength uint) float64 {
	hash := blake2s.Sum256(uintToBytesBE(data))
	return bytesToFloat64LE(hash[:], hashLength)
}

func Hash_md5(data, hashLength uint) float64 {
	hash := md5.Sum(uintToBytesLE(data))
	return bytesToFloat64LE(hash[:], hashLength)
}

func Hash_md4(data, hashLength uint) float64 {
	hash := md4.New().Sum(uintToBytesLE(data))
	return bytesToFloat64LE(hash[:], hashLength)
}

func Hash_bad(data, hashLength uint) float64 {
	bytes := uintToBytesLE(data)
	return bytesToFloat64LE(append(bytes[:1], []byte{0, 0, 0, 0, 0, 0, 0}...), hashLength)
}
