package stringMethods

type QNode struct {
	Next         *QNode
	Previous     *QNode
	Str          string
	NeedsClosure bool
}

func (qNode *QNode) reverse() {
	var newStr string = ""
	var startIndex int = len(qNode.Str) - 1
	for index := startIndex; index >= 0; index-- {
		var currentByte byte = qNode.Str[index]
		var currentStr string = string(currentByte)
		newStr = newStr + currentStr
	}
	qNode.Str = newStr
}

func FilterStringsToReverse(str string) string {
	var finalString string = ""
	var currentNode *QNode = &QNode{
		Next:         nil,
		Previous:     nil,
		Str:          "",
		NeedsClosure: false,
	}
	for index := range str {
		var currentByte byte = str[index]
		var currentString string = string(currentByte)
		if currentString != "(" && currentString != ")" {
			currentNode.Str = currentNode.Str + currentString
		}
		if currentString == "(" {
			if len(currentNode.Str) == 0 {
				currentNode.NeedsClosure = true
			} else {
				var newNode *QNode = &QNode{
					Next:         nil,
					Previous:     currentNode,
					Str:          "",
					NeedsClosure: true,
				}
				currentNode.Next = newNode
				currentNode = currentNode.Next
			}
		}

		if currentString == ")" {
			if currentNode.NeedsClosure {
				currentNode.reverse()
				if currentNode.Previous != nil {
					var reversedString string = currentNode.Str
					currentNode = currentNode.Previous
					currentNode.Str = currentNode.Str + reversedString
				}
			}
		}
	}
	for currentNode != nil {
		finalString = currentNode.Str + finalString
		currentNode = currentNode.Previous
	}
	return finalString
}
