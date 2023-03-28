package main

import (
	"fmt"
	"l2/pkg"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	Ns := make([]uint, 10000)

	for i := range Ns {
		Ns[i] = uint(i) + 1
	}

	ms := pkg.MultiSet_newMultiSet(4, 2)
	for elem, err := ms.Next(); err == nil; elem, err = ms.Next() {
		fmt.Println(elem)
	}

	// var wg sync.WaitGroup
	// wg.Add(5)
	// funcs := []func([]uint){pkg.Ex5a, pkg.Ex5b, pkg.Ex5c, pkg.Ex6, pkg.Ex7}

	// for _, f := range funcs {
	// 	ffunc := f
	// 	go func() {
	// 		defer wg.Done()
	// 		ffunc(Ns)
	// 	}()
	// }

	// wg.Wait()
	// pkg.Ex5b(Ns)
	// pkg.Ex5c(Ns) // 332
	pkg.Ex6(Ns)
	// pkg.ChGuysDeltas()
	// pkg.Ex7(Ns)
}
