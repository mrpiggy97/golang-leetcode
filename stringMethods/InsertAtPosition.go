package stringMethods

import "errors"

func InsertAtPosition(index int, slice []string, newMember string) []string {
	if index >= len(slice) {
		var newError error = errors.New("index cannot equal or be greater than length of slice")
		panic(newError)
	}

	var copySlice []string = []string{}
	copySlice = append(copySlice, slice[0:index]...)
	copySlice = append(copySlice, newMember)
	copySlice = append(copySlice, slice[index:]...)

	return copySlice
}
