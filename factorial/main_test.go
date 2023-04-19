package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFactorial_1to5(t1 *testing.T) {
	// arrange
	f40, ok := big.NewInt(0).SetString("815915283247897734345611269596115894272000000000", 10)
	if !ok {
		t1.Fatalf("initialisation of big int failed (f40)")
	}

	f120, ok := big.NewInt(0).SetString("6689502913449127057588118054090372586752746333138029810295671352301633557244962989366874165271984981308157637893214090552534408589408121859898481114389650005964960521256960000000000000000000000000000", 10)
	if !ok {
		t1.Fatalf("initialisation of big int failed (f120)")
	}

	tbl := []struct {
		n int64
		f *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(2)},
		{3, big.NewInt(6)},
		{4, big.NewInt(24)},
		{5, big.NewInt(120)},
		{40, f40},
		{120, f120},
	}

	for _, tt := range tbl {
		t1.Run(fmt.Sprintf("Test N=%v", tt.n), func(t *testing.T) {
			// act
			res := Factorial(big.NewInt(tt.n))

			// assert
			if res.Cmp(tt.f) != 0 {
				t.Errorf("Wrong result for %v!: %v, expected %v", tt.n, res, tt.f)
			}
		})
	}
}
