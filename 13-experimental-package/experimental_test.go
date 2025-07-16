package belajar_golang_generics

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.1, ExperimentalMin(100.1, 200.1))
}

func TestExperimentalMap(t *testing.T) {
	first := map[string]string{
		"Name": "Nafi",
	}

	second := map[string]string{
		"Name": "Nafi",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlice(t *testing.T) {
	first := []string{"Nafi"}
	second := []string{"Nafi"}
	assert.True(t, slices.Equal(first, second))
}
