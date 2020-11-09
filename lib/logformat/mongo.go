package logformat

import "fmt"

// MongoExecuteError default mongo response error message
func MongoExecuteError(err error) error {
	return fmt.Errorf("mongodb error: %v", err)
}

// MongoExecuteWithValueError default mongo response error message
func MongoExecuteWithValueError(value interface{}, err error) error {
	return fmt.Errorf("mongodb error with parameter %v: %v", value, err)
}
