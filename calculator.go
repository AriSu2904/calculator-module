package calculator_module

import "errors"

func Sum(a int, b int) int {
	return a + b
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa dibagi dengan 0!")
	}
	return a / b, nil
}
