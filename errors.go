package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil

}

func firstFunction() error {
	return fmt.Errorf("original error : something went wrong in FirstFunction")
}
func secondfunc() error {
	firstErr := firstFunction()
	if firstErr != nil {
		secondErr := errors.New("failed in second Function")
		return errors.Join(firstErr, secondErr)
	}
	return nil

}
func errorkaro() {
	// ans, err := divide(10, 2)
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return
	// }
	// fmt.Println(ans)

	err := secondfunc()
	fmt.Println("original Error", err)
	innerError := errors.Unwrap(err)
	fmt.Println("Inner Error", innerError)

}
