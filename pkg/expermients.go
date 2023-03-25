package pkg

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Ex5b(Ns []uint) {
	fmt.Println("Running Ex5b")

	ks := []uint{2, 3, 10, 100, 400}
	filename := "data/exp5b.txt"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer f.Close()

	for _, k := range ks {
		resForK := make([]string, len(Ns))
		for i, n := range Ns {
			ms := MultiSet_newMultiSet(n)
			count := Mincount(ms, Hash_blake2b, k, 6)
			resForK[i] = fmt.Sprintf("%v", count/float64(n))

			// if i%1000 == 0 {
			// 	fmt.Printf("Done n = %v for k = %v\n", n, k)
			// }
		}

		fmt.Fprint(f, strings.Join(resForK, " ")+"\n")
		fmt.Println("Done Ex5b for k = ", k)
	}

	fmt.Println("Done Ex5b")

}

func Ex5c(Ns []uint) {
	fmt.Println("Running Ex5c")
	expectedGood := uint(0.95 * float64(len(Ns)))

	var left, right uint = 2, 400

	for left < right {
		var mid = (left + right) / 2
		var goodCounter uint = 0

		fmt.Println("Running Ex5c for k = ", mid)

		res := runForGivenK(Ns, mid, Hash_blake2b, 6)

		for _, n := range res {
			if n > 0.9 && n < 1.1 {
				goodCounter++
			}
		}

		if goodCounter >= expectedGood {
			right = mid
		} else {
			left = mid + 1
		}

	}
	fmt.Println("Best k = ", left)
	fmt.Println("Done Ex5c")
}

func Exp6(Ns []uint) {
	fmt.Println("Running Ex6")

	bytes := []uint{1, 2, 3, 4, 5, 6}

	for _, b := range bytes {
		fmt.Println("Running for b = ", b)

		avgSha1 := getAvgRes(Ns, 333, Hash_sha1, b)
		avgSha256 := getAvgRes(Ns, 333, Hash_sha256, b)
		avgBlake2 := getAvgRes(Ns, 333, Hash_blake2b, b)

		fmt.Printf("Sha1 avg difference = %G\n", avgSha1)
		fmt.Printf("Sha256 avg difference = %G\n", avgSha256)
		fmt.Printf("ShaBlake2b avg difference = %G\n", avgBlake2)
	}

	fmt.Println("Done Ex6")
}

func runForGivenK(Ns []uint, k uint, h func(uint, uint) float64, hashLength uint) []float64 {
	res := make([]float64, len(Ns))
	for i, n := range Ns {
		ms := MultiSet_newMultiSet(n)
		count := Mincount(ms, h, k, hashLength)
		res[i] = count / float64(n)
	}

	return res
}

func getAvgRes(Ns []uint, k uint, h func(uint, uint) float64, hashLength uint) float64 {
	res := runForGivenK(Ns, k, h, hashLength)

	sumOfDiffs := 0.0

	for _, f := range res {
		sumOfDiffs = math.Abs(f - 1.0)
	}

	return sumOfDiffs

}
