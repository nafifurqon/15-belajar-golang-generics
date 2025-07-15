package belajar_golang_generics

import "testing"

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Muhammad", "Nafi", "Furqon"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(numbers)
}
