package ex2

import (
	"euler/app"
	"euler/exercies"
	"math"
	"strconv"
)

const (
	Name     = "2"
	PathTest = "./exercies/ex2/test.txt"
)

//Even Fibonacci Numbers
/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
By considering the terms in the Fibonacci sequence whose values do not exceed n, find the sum of the even-valued terms.
*/
func Ex(number int64) uint64 {

	var (
		sum = uint64(0)
		a   = int64(1)
		b   = int64(2)
		nr  = a + b
	)
	if number > 2 {
		sum += 2
	}

	for nr < number {
		a = b
		b = nr
		if nr%2 == 0 {
			sum += uint64(nr)
		}

		nr = a + b

	}

	return sum
}

func NthFibonnaciNumber(number int64) uint64 {
	sqrt5 := math.Sqrt(5)
	x := 1 / sqrt5
	y := math.Pow((1+sqrt5), float64(number)) / math.Pow(float64(2), float64(number))

	return uint64(math.Floor(x * y))
}

func Register() {

	//Exercies
	app.AddExercies(Name, func(inputs ...string) []interface{} {
		if len(inputs) < 1 {
			panic("Expected one parameter of type int")
		}

		nr, err := strconv.ParseInt(inputs[0], 10, 64)

		if err != nil {
			panic(err.Error())
		}

		return []interface{}{
			Ex(nr),
		}
	})

	//Tests
	app.AddTest(Name, exercies.EasyTest(Name, PathTest))
}
