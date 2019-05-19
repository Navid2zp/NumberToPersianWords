package NumberToPersianWords

import (
	"math"
)

func NumberToWords(number int) string {
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
	return Ones[number]
}

func twoDigits(number int) string {
	if number < 20 {
		return TenToTwenty[(number % 10)]
	} else if number%10 == 0 {
		return Tens[(number/10)-2]
	}
	return Tens[((number-(number%10))/10)-2] + " و " + oneDigit(number%10)
}

func threeDigits(number int) string {
	if number%100 == 0 {
		return Hundreds[(number/100) - 1]
	}
	return Hundreds[(number/100)-1] + " و " + NumberToWords(number%100)
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
		return NumberToWords(newValue) + " " + Classes[counter] + " و " + NumberToWords(leftOver)
	}
	return NumberToWords(newValue) + " " + Classes[counter]
}
