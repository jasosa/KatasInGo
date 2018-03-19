package fizzbuzz

import "math"
import "strconv"

func GetFizzBuzz(number int) string {

	var fizzbuzz = ""

	if math.Mod(float64(number), 3) != 0 && math.Mod(float64(number), 5) != 0 {
		fizzbuzz = strconv.Itoa(number)
	} else {

		if math.Mod(float64(number), 3) == 0 {
			fizzbuzz += "Fizz"
		}

		if math.Mod(float64(number), 5) == 0 {
			fizzbuzz += "Buzz"
		}
	}

	return fizzbuzz
}
