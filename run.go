package NumberToPersianWords

import (
	"math"
)

func ParseInt(number int) string {
	switch {
	case number < 10:
		return oneDigit(number)
	case number < 100:
		return twoDigits(number)
	case number < 1000:
		return threeDigits(number)
	default:
		return bigOnes(number)
	}
}

func oneDigit(number int) string {
	return ones[number]
}

func twoDigits(number int) string {
	if number < 20 {
		return tenToTwenty[(number % 10)]
	} else if number%10 == 0 {
		return tens[(number/10)-2]
	}
	return tens[(number/10)-2] + " و " + oneDigit(number%10)
}

func threeDigits(number int) string {
	if number%100 == 0 {
		return hundreds[(number/100) - 1]
	}
	return hundreds[(number/100)-1] + " و " + ParseInt(number%100)
}

func bigOnes(number int) string {
	counter := -1
	newValue := number
	for newValue > 999 {
		newValue /= 1000
		counter += 1
	}

	leftOver := number - (int(math.Pow(1000.0, float64(counter + 1))) * newValue)
	if leftOver > 0 {
		return ParseInt(newValue) + " " + classes[counter] + " و " + ParseInt(leftOver)
	}
	return ParseInt(newValue) + " " + classes[counter]
}
