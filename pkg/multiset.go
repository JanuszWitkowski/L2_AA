package pkg

import "errors"

type MultiSet struct {
	n         uint
	counter   uint
	first     uint
	elemsLeft uint
}

func MultiSet_newMultiSet(n uint, nums ...uint) *MultiSet {
	var left uint = n

	if len(nums) > 0 {
		left = nums[0] * left
	}

	return &MultiSet{
		n:         n,
		counter:   0,
		first:     (n * (n - 1)) / 2,
		elemsLeft: left,
	}
}

func (ms *MultiSet) Next() (uint, error) {
	if ms.elemsLeft == 0 {
		return 0, errors.New("no elems left")
	}

	ret := ms.first + ms.counter

	ms.elemsLeft -= 1
	ms.counter = (ms.counter % ms.n) + 1

	return ret, nil
}

// n = 5
// first == 10
// 11, 12, 13, 14, 15, 11

// n = 4
// first == 6
// 7, 8, 9, 10, 7, 8, 9, 10
