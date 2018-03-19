package fizzbuzz

import "testing"

var fizzbuzztests = []struct {
	in  int
	out string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{9, "Fizz"},
	{11, "11"},
	{15, "FizzBuzz"},
	{17, "17"},
	{20, "Buzz"},
	{21, "Fizz"},
	{23, "23"},
	{30, "FizzBuzz"},
}

//TestFizzbuzz tests...
func TestFizzbuzz(t *testing.T) {
	for _, tt := range fizzbuzztests {
		result := GetFizzBuzz(tt.in)
		if tt.out != result {
			t.Errorf("Input: %d Expected: %s Received: %s", tt.in, tt.out, result)
		}
	}
}
