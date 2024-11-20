package main

import "errors"

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func Multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

var errDivisaoPorZero = errors.New("divisão por zero")

func Divide(i ...float64) (float64, error) {

	if len(i) == 0 {
		return 0, nil
	}
	if len(i) == 1 {
		return i[0], nil
	}
	total := i[0]
	for index, v := range i {
		if index == 0 {
			continue
		}
		if v == 0 {
			return 0, errDivisaoPorZero
		}
		total /= v
	}
	return total, nil
}

func Subtrai(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	if len(i) == 1 {
		return i[0]
	}
	total := i[0]
	for index, v := range i {
		if index == 0 {
			continue
		}
		total -= v
	}
	return total
}

func main() {

}
