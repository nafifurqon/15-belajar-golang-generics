package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(param T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}

	result := ChangeValue(&myData, "Nafi")

	assert.Equal(t, "Nafi", result)
	assert.Equal(t, "Nafi", myData.Value)
}
