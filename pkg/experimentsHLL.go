package pkg

import (
	"fmt"
	"os"
	"strings"
)

func Ex5aHLL(Ns []uint) {
	fmt.Println("Running Ex5a")

	Ms := []uint{1, 2, 3, 4}

	filename := "data/exp5aHLL.txt"
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
			count := Hyperloglog(ms, Hash_blake2b_PURE, 10)
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

func Ex5bHLL(Ns []uint) {
	fmt.Println("Running Ex5b")

	ks := []uint32{2, 5, 7, 10, 16}
	filename := "data/exp5bHLL.txt"
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
			count := Hyperloglog(ms, Hash_blake2b_PURE, k)
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

func Ex5cHLL(Ns []uint) {
	fmt.Println("Running Ex5c")
	expectedGood := uint(0.95 * float64(len(Ns)))

	var left, right uint32 = 1, 32

	for left < right {
		var mid = (left + right) / 2
		var goodCounter uint = 0

		fmt.Println("Running Ex5c for k = ", mid)

		res := runForGivenBHLL(Ns, Hash_blake2b_PURE, mid)

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

func Compare(Ns []uint) {
	fmt.Println("Running Compare")
	filename := "data/compare.txt"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer f.Close()

	resHLL := runForGivenBHLL(Ns, Hash_blake2b_PURE, 5)
	resMC := runForGivenKMC(Ns, 5, Hash_blake2b, 4)

	for i := range resHLL {
		fmt.Fprintf(f, "%f %f\n", resHLL[i], resMC[i])

	}

	fmt.Println("Done Compare")
}

func runForGivenBHLL(Ns []uint, h func(uint, uint) []byte, b uint32) []float64 {
	res := make([]float64, len(Ns))
	for i, n := range Ns {
		ms := MultiSet_newMultiSet(n)
		count := Hyperloglog(ms, h, b)
		res[i] = count / float64(n)
	}

	return res
}
