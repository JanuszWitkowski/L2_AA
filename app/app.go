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
	// funcs := []func([]uint){pkg.Ex5aHLL, pkg.Ex5bHLL, pkg.Ex5cHLL}

	// for _, f := range funcs {
	// 	ffunc := f
	// 	go func() {
	// 		defer wg.Done()
	// 		ffunc(Ns)
	// 	}()
	// }

	// wg.Wait()
	// pkg.Ex5aHLL(Ns)
	// pkg.Ex5bHLL(Ns)
	pkg.Ex5cHLL(Ns) // 332
	// pkg.Ex6(Ns)
	// pkg.ChGuysDeltas()
	// pkg.Ex7(Ns)
}
