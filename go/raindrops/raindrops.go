package raindrops

import "strconv"

func Convert(num int) string  {
	convertedString := ""

	if num % 3 == 0 {
		convertedString += "Pling"
	}

	if num % 5 == 0 {
		convertedString += "Plang"
	}

	if num % 7 == 0 {
		convertedString += "Plong"
	}

	if convertedString == "" {
		convertedString = strconv.Itoa(num)
	}

	return convertedString
}

