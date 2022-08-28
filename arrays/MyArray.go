package arrays

import "fmt"

type Array[T comparable] struct {
	Data   map[int]*T
	Length *int
}

func (array *Array[T]) Append(value *T) {
	array.Data[*array.Length] = value
	*array.Length = *array.Length + 1
}

func (array *Array[T]) PrintMembers() {
	for index := 0; index < *array.Length; index++ {
		fmt.Println(*array.Data[index])
	}
}

func (array *Array[T]) GetIndexOfValue(value T) *int {
	var correctIndex *int = nil
	for index := 0; index < *array.Length; index++ {
		var element T = *array.Get(index)
		if element == value {
			fmt.Println("element and index", index, element)
			correctIndex = &index
			break
		}
	}

	return correctIndex
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

func (array *Array[T]) InsertAt(index int, value *T) {
	if index >= *array.Length {
		fmt.Println("index out of bounds for value ", *value)
	} else {
		var valuesToMove []*T = []*T{}
		// append all elements that from index foward to valuesToMove
		for currentIndex := index; currentIndex < *array.Length; currentIndex++ {
			valuesToMove = append(valuesToMove, array.Get(currentIndex))
		}
		// replace item in index with new value
		array.Data[index] = value
		// increase length of array
		*array.Length = *array.Length + 1
		// index of array in valuesToMove
		var indexOfElement int = 0
		// loop through every element after the new element inserted
		// replace the current value of that element with elements
		// in valuesToMove slice
		for currentIndex := index + 1; currentIndex < *array.Length; currentIndex++ {
			array.Data[currentIndex] = valuesToMove[indexOfElement]
			indexOfElement = indexOfElement + 1
		}
	}
}

func (array *Array[T]) GetSlice() []T {
	if *array.Length == 0 {
		return nil
	}
	var values []T = []T{}
	for index := 0; index < *array.Length; index++ {
		values = append(values, *array.Get(index))
	}
	return values
}

func NewArray[T comparable]() *Array[T] {
	var length int = 0
	return &Array[T]{
		Data:   make(map[int]*T),
		Length: &length,
	}
}
