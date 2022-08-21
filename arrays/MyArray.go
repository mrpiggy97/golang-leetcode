package arrays

import "fmt"

type Array[T any] struct {
	Data   map[int]*T
	Length *int
}

func (array *Array[T]) Append(value *T) {
	array.Data[*array.Length] = value
	*array.Length = *array.Length + 1
}

func (array *Array[T]) PrintMembers() {
	for index := 0; index < *array.Length; index++ {
		fmt.Println(array.Data[index])
	}
}

func (array *Array[T]) Get(index int) *T {
	valueToReturn, exists := array.Data[index]
	if exists {
		return valueToReturn
	} else {
		return nil
	}
}

func (array Array[T]) InsertAtStart(value *T) {
	if *array.Length < 0 {
		fmt.Println("array can't have a length of zero")
	}

	if *array.Length > 0 {
		*array.Length = *array.Length + 1
		var currentValue *T = array.Get(0)
		for currentIndex := 0; currentIndex < *array.Length; currentIndex++ {
			if currentIndex == 0 {
				array.Data[currentIndex] = value
			}
			if currentIndex > 0 {
				temporaryValue, exists := array.Data[currentIndex]
				array.Data[currentIndex] = currentValue
				if exists {
					currentValue = temporaryValue
				}
			}
		}
	}
}

func NewArray[T any]() *Array[T] {
	var length int = 0
	return &Array[T]{
		Data:   make(map[int]*T),
		Length: &length,
	}
}
