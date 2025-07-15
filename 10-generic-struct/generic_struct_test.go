package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

// error: method must have no type parameters
// func (d *Data[T]) ChangeFirst[E any](first T) {}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	// error: cannot use generic type Data[T any] without instantiation
	// tidak support type inference
	// data := Data{
	// 	First:  "Nafi",
	// 	Second: "Furqon",
	// }

	data := Data[string]{
		First:  "Nafi",
		Second: "Furqon",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Nafi",
		Second: "Furqon",
	}

	assert.Equal(t, "Muhammad", data.ChangeFirst("Muhammad"))
	assert.Equal(t, "Hello Diani", data.SayHello("Diani"))
}
