package stringMethods

type QNode struct {
	Previous *QNode
	Str      string
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
	var currentStr string = ""
	var currentNode *QNode = nil
	for index := range str {
		var currentByte byte = str[index]
		var stringValue string = string(currentByte)
		if stringValue != "(" && stringValue != ")" {
			currentStr = currentStr + stringValue
		}
		if stringValue == "(" && len(currentStr) > 0 {
			if currentNode == nil {
				currentNode = &QNode{
					Previous: nil,
					Str:      currentStr,
				}
			} else {
				var newNode *QNode = &QNode{
					Previous: currentNode,
					Str:      currentStr,
				}
				currentNode = newNode
			}
			currentStr = ""
		}
		if stringValue == ")" {
			if len(currentStr) > 0 {
				if currentNode == nil {
					currentNode = &QNode{
						Previous: nil,
						Str:      currentStr,
					}
				} else {
					var newNode *QNode = &QNode{
						Previous: currentNode,
						Str:      currentStr,
					}
					currentNode = newNode
				}
				currentNode.reverse()
				if currentNode.Previous != nil {
					var reversedString string = currentNode.Str
					currentNode = currentNode.Previous
					currentNode.Str = currentNode.Str + reversedString
				}
				currentStr = ""
			} else {
				if currentNode != nil {
					currentNode.reverse()
					var reversedString string = currentNode.Str
					currentNode = currentNode.Previous
					currentNode.Str = currentNode.Str + reversedString
				} else {
					return ""
				}
			}
		}
	}
	if len(currentStr) > 0 {
		var newNode *QNode = &QNode{
			Previous: currentNode,
			Str:      currentStr,
		}
		currentNode = newNode
	}
	var finalString string = ""
	for currentNode != nil {
		finalString = currentNode.Str + finalString
		currentNode = currentNode.Previous
	}
	return finalString
}
