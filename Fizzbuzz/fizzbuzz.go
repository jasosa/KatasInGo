package main

import "fmt"
import "math"
import "strconv"

func getFizzBuzz(number int) (string) {

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


func main(){
	for index := 1 ; index < 101 ; index++{	
		fmt.Println(getFizzBuzz(index));
	}
}

