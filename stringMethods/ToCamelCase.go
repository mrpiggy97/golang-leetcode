package stringMethods

import (
	"strings"
	"time"
	"unicode"
)

func ToCamelCase(stringInstance string, channel chan<- string) {
	time.Sleep(time.Second * 2)
	var newString string = ""
	var stringSlice []string = strings.Split(stringInstance, " ")
	for index, letter := range stringSlice {
		var byteSlice []byte = []byte(letter)
		for idx, byteLetter := range byteSlice {
			if idx == 0 {
				byteSlice[idx] = byte(unicode.ToUpper(rune(byteLetter)))
			} else {
				byteSlice[idx] = byte(unicode.ToLower(rune(byteLetter)))
			}
		}

		newString = newString + string(byteSlice)
		if index < len(stringSlice)-1 {
			newString = newString + " "
		}
	}
	channel <- newString
}

func GetConvertedString(stringToConvert string, channel chan string) string {
	go ToCamelCase(stringToConvert, channel)
	stringToConvert = <-channel
	return stringToConvert
}
