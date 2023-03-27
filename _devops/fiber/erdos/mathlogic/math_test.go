package mathlogic

import (
	"math/big"
	"testing"
)

type data struct {
	x, y   int64
	result int64
}

type Bigdata struct {
	x   *big.Int
	y   *big.Int
	result *big.Int
}

func TestAdd(t *testing.T) {

	testData := []data{
		{993, 2, 995},
		{543, 7, 550},
		{321, -321, 0},
		{-321, -321, -642},
	}

	for _, datum := range testData {
		result := Add(datum.x, datum.y)
		if result != datum.result {
			t.Errorf("Add(%d, %d) FAILED. Expected %d got %d\n",
				datum.x, datum.y, datum.result, result)
		}
	}
}

func TestSub(t *testing.T) {
	var expected1, _ = new(big.Int).SetString("0", 0)
	var expected2, _ = new(big.Int).SetString("-95164610537311803229070505640454897895348941329229195312512972132373216059", 0)
	var expected3, _ = new(big.Int).SetString("95164610537311803229070505640454897895348941329229195312512972132373216059", 0)
	var bigX, _ = new(big.Int).SetString("-8651328230664709384460955058223172535940812848111745028410270193852110500", 0)
	var bigY, _ = new(big.Int).SetString("86513282306647093844609550582231725359408128481117450284102701938521105559", 0)

	testData := []Bigdata{
		{bigY, bigY, expected1},
		{bigX, bigY, expected2},
		{bigY, bigX, expected3},
	}

	for _, datum := range testData {
		result := Sub(datum.x, datum.y)
		if result.Cmp(datum.result) !=0 {
			t.Errorf("Sub(%d, %d) FAILED. Expected %d got %d\n",
				datum.x, datum.y, datum.result, result)
		}
	}
}

func TestMul(t *testing.T) {

	var expected, _ = new(big.Int).SetString("122496452125610704938685320429399871474856557930892133859113727985847608675495662329857420238167604161402916835750863687571688881049043110684664170679578461660476299571022355607320", 0)
	var bigX, _ = new(big.Int).SetString("1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480", 0)
	var bigY, _ = new(big.Int).SetString("86513282306647093844609550582231725359408128481117450284102701938521105559", 0)

	result := Mul(bigX, bigY)

	if result.Cmp(expected) != 0 {
		t.Errorf("Mul(%d %d) FAILED. Expected %d got %d\n", bigX, bigY, expected, result)
	}

}

func TestDivide(t *testing.T) {

	result := Divide(5, 0)

	if result != 0 {
		t.Errorf("Divide(5, 0) FAILED. Expected %f, got %f\n", 0.0, result)
	} else {
		t.Logf("Divide(5, 0) PASSED. Expected %f, got %f\n", 0.0, result)
	}
}


func TestAddBase(t *testing.T) {
	testCases := []struct {
		x, y     int64
		expected int64
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, tc := range testCases {
		result := Add(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", tc.x, tc.y, result, tc.expected)
		}
	}
}

func TestSubBase(t *testing.T) {
	testCases := []struct {
		x, y     *big.Int
		expected string
	}{
		{big.NewInt(10), big.NewInt(5), "5"},
		{big.NewInt(-10), big.NewInt(-5), "-5"},
		{big.NewInt(0), big.NewInt(0), "0"},
	}

	for _, tc := range testCases {
		result := Sub(tc.x, tc.y)
		if result.String() != tc.expected {
			t.Errorf("Sub(%s, %s) = %s; expected %s", tc.x.String(), tc.y.String(), result.String(), tc.expected)
		}
	}
}

func TestDivideBase(t *testing.T) {
	testCases := []struct {
		x, y     int64
		expected float64
	}{
		{10, 5, 2},
		{0, 5, 0},
		{10, 0, 0},
	}

	for _, tc := range testCases {
		result := Divide(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("Divide(%d, %d) = %f; expected %f", tc.x, tc.y, result, tc.expected)
		}
	}
}

func TestMulBase(t *testing.T) {
	testCases := []struct {
		x, y     *big.Int
		expected string
	}{
		{big.NewInt(10), big.NewInt(5), "50"},
		{big.NewInt(-10), big.NewInt(-5), "50"},
		{big.NewInt(0), big.NewInt(0), "0"},
	}

	for _, tc := range testCases {
		result := Mul(tc.x, tc.y)
		if result.String() != tc.expected {
			t.Errorf("Mul(%s, %s) = %s; expected %s", tc.x.String(), tc.y.String(), result.String(), tc.expected)
		}
	}
}
