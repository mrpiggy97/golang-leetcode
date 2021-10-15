package arithmetic

func FindMembers(members []bool) int {
	var counter int = 0
	for _, member := range members {
		if member {
			counter++
		}
	}

	return counter
}
