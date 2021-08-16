package korean_romanizer_go

func IndexRune(array []rune, char rune) int {
	for i, runeValue := range array {
		if char == runeValue {
			return i
		}
	}
	return -1

}
