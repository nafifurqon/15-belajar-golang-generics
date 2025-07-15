package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

type NumberApproximation interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func MinApproximation[T NumberApproximation](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100), Min[float64](float64(100), float64(200)))
	// assert.Equal(t, float64(100), Min[float64](float64(100), int64(200))) // error: cannot use int64(200) (constant 200 of type int64) as float64 value in argument to Min[float64]
	// assert.Equal(t, Age(100), Min[Age](Age(100), Age(200))) // error: Age does not satisfy Number (possibly missing ~ for int in Number)
}

func TestMinApproxiamtion(t *testing.T) {
	assert.Equal(t, 100, MinApproximation[int](100, 200))
	assert.Equal(t, int64(100), MinApproximation[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100), MinApproximation[float64](float64(100), float64(200)))
	assert.Equal(t, Age(100), MinApproximation[Age](Age(100), Age(200))) // error: Age does not satisfy Number (possibly missing ~ for int in Number)
}

func TestMinTypeInference(t *testing.T) {
	assert.Equal(t, 100, MinApproximation(100, 200))
	assert.Equal(t, int64(100), MinApproximation(int64(100), int64(200)))
	assert.Equal(t, float64(100), MinApproximation(float64(100), float64(200)))
	assert.Equal(t, Age(100), MinApproximation(Age(100), Age(200))) // error: Age does not satisfy Number (possibly missing ~ for int in Number)
}
