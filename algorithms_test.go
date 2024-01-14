package algorithms_test

import (
	algorithms "algorithms/numbers"
	"fmt"
	"testing"
)

func TestFindFactors(t *testing.T) {
	factors := algorithms.FindFactors(568)
	fmt.Println(factors)
}
