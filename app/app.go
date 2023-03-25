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

	// ms := pkg.MultiSet_newMultiSet(5)
	// for elem, err := ms.Next(); err == nil; elem, err = ms.Next() {
	// 	fmt.Println(elem)
	// }
	// pkg.Ex5b(Ns)
	// pkg.Ex5c(Ns)
	pkg.Exp6(Ns)
}
