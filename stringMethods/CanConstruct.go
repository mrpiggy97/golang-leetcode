package stringMethods

import (
	"fmt"
	"strings"
)

func printEndMaps(map1, map2 map[string]int) {
	fmt.Println("map1 ", map1)
	fmt.Println("map2 ", map2)
}

func CanConstruct(randomNote string, magazine string) bool {
	var randomNoteSlice []string = strings.Split(randomNote, "")
	var magazineSlice []string = strings.Split(magazine, "")
	var randomNoteMap map[string]int = make(map[string]int)
	var magazineMap map[string]int = make(map[string]int)
	defer printEndMaps(randomNoteMap, magazineMap)
	for _, letter := range randomNoteSlice {
		var count int = strings.Count(randomNote, letter)
		randomNoteMap[letter] = count
	}

	for _, letter := range magazineSlice {
		if strings.Contains(randomNote, letter) {
			var count int = strings.Count(magazine, letter)
			magazineMap[letter] = count
		}
	}

	if len(randomNoteMap) != len(magazineMap) {
		return false
	}

	for letter, count := range randomNoteMap {
		var magazineCount int = magazineMap[letter]
		if magazineCount < count {
			return false
		}
	}
	return true
}
