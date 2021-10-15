package arithmetic

func MultiplyAllMembers(members []int) int {
	var currentValue int = members[0] * members[1]
	for _, nextNumber := range members[2:] {
		currentValue = currentValue * nextNumber
	}

	return currentValue
}
