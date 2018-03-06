package main

import "fmt"
import "math"
import "strconv"

func getFizzBuzz(number int) (string) {

	var fizzbuzz = ""

	/*if math.Mod(float64(number), 3) == 0 && math.Mod(float64(number), 5) == 0 {
		fizzbuzz = "FizzBuzz"
	} else if math.Mod(float64(number), 5) == 0 {
		fizzbuzz = "Buzz"
	} else if math.Mod(float64(number), 3) == 0 {
		fizzbuzz = "Fizz"
	} else {
		fizzbuzz = strconv.Itoa(number)
	}*/

	if math.Mod(float64(number), 3) != 0 && math.Mod(float64(number), 5) != 0{
		fizzbuzz = strconv.Itoa(number)
	}else{

		if math.Mod(float64(number), 3) == 0 {
			fizzbuzz += "Fizz"
		}

		if math.Mod(float64(number), 5) == 0 {
			fizzbuzz += "Buzz"
		}
	}

	return fizzbuzz
}


func main(){
	for index := 0 ; index < 100 ; index++{	
		fmt.Println(getFizzBuzz(index));
	}
}

