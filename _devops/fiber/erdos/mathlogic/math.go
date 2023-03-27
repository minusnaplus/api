package mathlogic

import "math/big"

func Add(x, y int64) int64 {
	return x + y
}

func Sub(x, y *big.Int) *big.Int {
    return big.NewInt(0).Sub(x, y)
}

func Divide(x, y int64) float64 {
	if y == 0 {
		return float64(0)
	}
	return float64(x / y)
}

func Mul(x, y *big.Int) *big.Int {
    return big.NewInt(0).Mul(x, y)
}

