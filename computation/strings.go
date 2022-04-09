package computation

import "strconv"

func SplitString(dateString string, symbol string) []int {

	var stringArray [3]string
	temp := ""
	i := 0
	for _, character := range dateString {
		if character == '/' {
			stringArray[i] = temp
			i++
			temp = ""
		} else {
			temp += string(character)
		}
	}
	stringArray[2] = temp

	integerArray := make([]int, 0, 3)
	for _, str := range stringArray {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		integerArray = append(integerArray, i)
	}

	return integerArray
}
