package computation

import "strconv"

//SplitString - will split the string
func SplitString(dateString string) []int {

	var stringArray [3]string
	temp := ""
	i := 0

	//going through each character in a string
	for _, character := range dateString {

		//if block seperates the seperator and string
		if character == SeperatorRune {
			stringArray[i] = temp
			i++
			temp = ""
		} else {
			temp += string(character)
		}
	}
	stringArray[2] = temp

	//converts the strings into integer and returns it
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
