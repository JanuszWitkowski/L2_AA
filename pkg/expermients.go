package pkg

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func Ex5a(Ns []uint) {
	fmt.Println("Running Ex5a")

	Ms := []uint{1, 2, 3, 4}

	filename := "data/exp5a.txt"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer f.Close()

	for _, m := range Ms {
		resForM := make([]string, len(Ns))
		for i, n := range Ns {
			ms := MultiSet_newMultiSet(n, m)
			count := Mincount(ms, Hash_blake2b, 333, 6)
			resForM[i] = fmt.Sprintf("%v", count/float64(n))

			// if i%1000 == 0 {
			// 	fmt.Printf("Done n = %v for k = %v\n", n, k)
			// }
		}

		fmt.Fprint(f, strings.Join(resForM, " ")+"\n")
		fmt.Println("Done Ex5a for m = ", m)
	}
	fmt.Println("Done Ex5a")
}

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

		res := runForGivenKMC(Ns, mid, Hash_blake2b, 6)

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

func Ex6(Ns []uint) {
	fmt.Println("Running Ex6")

	// filename := "data/exp6.txt"
	// f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	// defer f.Close()

	bytes := []uint{1, 2, 3, 4, 5, 6}

	for _, b := range bytes {
		fmt.Println("Running for b = ", b)

		// avgSha1 := getAvgRes(Ns, 333, Hash_sha1, b)
		// avgSha3 := getAvgRes(Ns, 333, Hash_sha3, b)
		// avgSha256 := getAvgRes(Ns, 333, Hash_sha256, b)
		// avgBlake2b := getAvgRes(Ns, 333, Hash_blake2b, b)
		// avgBlake2s := getAvgRes(Ns, 333, Hash_blake2s, b)
		// avgMd5 := getAvgRes(Ns, 333, Hash_md5, b)
		// avgMd4 := getAvgRes(Ns, 333, Hash_md4, b)
		avgBad := getAvgRes(Ns, 333, Hash_bad, b)

		// fmt.Printf("Sha1 avg difference = %G\n", avgSha1)
		// fmt.Printf("Sha3 avg difference = %G\n", avgSha3)
		// fmt.Printf("Sha256 avg difference = %G\n", avgSha256)
		// fmt.Printf("ShaBlake2b avg difference = %G\n", avgBlake2b)
		// fmt.Printf("ShaBlake2s avg difference = %G\n", avgBlake2s)
		// fmt.Printf("Md5 avg difference = %G\n", avgMd5)
		// fmt.Printf("Md4 avg difference = %G\n", avgMd4)
		fmt.Printf("Bad avg difference = %G\n", avgBad)

		// fmt.Fprintf(f, "%G %G %G %G %G %G %G %G\n", avgSha1, avgSha3, avgSha256, avgBlake2b, avgBlake2s, avgMd5, avgMd4, avgBad)
	}

	fmt.Println("Done Ex6")
}

func Ex7(Ns []uint) {
	fmt.Println("Running Ex7")
	var k uint = 400
	alphas := []uint{9500, 9900, 9950}
	h := Hash_sha256

	res := runForGivenKMC(Ns, k, h, 4)

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })

	for _, alpha := range alphas {
		right := 1.0
		left := 0.0
		fmt.Println("Running Ex5c for alpha = ", alpha)
		for left <= right {
			var mid = (left + right) / 2

			counter := getCountInRangeInSorted(res, mid)

			if counter >= alpha {
				right = math.Nextafter(mid, 0.0)
			} else {
				left = math.Nextafter(mid, math.MaxFloat64)
			}

		}
		fmt.Println("Delta for alpha = ", alpha, " - ", left)
	}

	fmt.Println("Done Ex7")
}

func runForGivenKMC(Ns []uint, k uint, h func(uint, uint) float64, hashLength uint) []float64 {
	res := make([]float64, len(Ns))
	for i, n := range Ns {
		ms := MultiSet_newMultiSet(n)
		count := Mincount(ms, h, k, hashLength)
		res[i] = count / float64(n)
	}

	return res
}

func getAvgRes(Ns []uint, k uint, h func(uint, uint) float64, hashLength uint) float64 {
	res := runForGivenKMC(Ns, k, h, hashLength)

	sumOfDiffs := 0.0

	for _, f := range res {
		sumOfDiffs = math.Abs(f - 1.0)
	}

	return sumOfDiffs

}

func getCountInRangeInSorted(results []float64, delta float64) uint {
	low, high := 0, len(results)-1

	for i, f := range results {
		if f >= 1.0-delta {
			low = i
			break
		}
	}

	for i := high; i >= 0; i-- {
		if results[i] <= 1.0+delta {
			high = i
			break
		}
	}

	return uint(high-low) + 1
}

func chebyschevHelper(alpha float64, k uint) {
	fmt.Println("Chebyschev for alpha = ", alpha, " delta = ", math.Sqrt(1.0/(float64(k)*alpha)))
}

func chernoffHelper(delta float64, k uint) float64 {
	eps1 := delta / (1 - delta)
	eps2 := delta / (1 + delta)

	fk := func(arg float64) float64 {
		return math.Exp(arg*float64(k)) * math.Pow(1.0-arg, float64(k))
	}

	return fk(eps2) + fk(-eps1)
}

func chernoffBinSearch(alpha float64, k uint) {

	left := 0.0
	right := 1.0

	for left <= right {
		var mid = (left + right) / 2

		res := chernoffHelper(mid, k)
		if res <= alpha {
			right = math.Nextafter(mid, 0.0)
		} else {
			left = math.Nextafter(mid, math.MaxFloat64)
		}

	}

	fmt.Println("Chernoff For alpha = ", alpha, " delta =", left)

}

func ChGuysDeltas() {
	alphas := []float64{0.05, 0.01, 0.005}
	var k uint = 400

	for _, a := range alphas {
		chebyschevHelper(a, k)
		chernoffBinSearch(a, k)
	}
}
