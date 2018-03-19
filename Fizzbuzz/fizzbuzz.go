package fizzbuzz

import "math"
import "strconv"

//GetFizzBuzz sdfdsf
func GetFizzBuzz(number int) string {

	var fizzbuzz = ""

	if math.Mod(float64(number), 3) == 0 {
		fizzbuzz += "Fizz"
	}

	if math.Mod(float64(number), 5) == 0 {
		fizzbuzz += "Buzz"
	}

	if fizzbuzz == "" {
		fizzbuzz = strconv.Itoa(number)
	}

	return fizzbuzz
}
