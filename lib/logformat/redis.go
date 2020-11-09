package logformat

import (
	"encoding/json"
	"fmt"
)

// RedisExecuteError default redis response error message
func RedisExecuteError(err error) error {
	return fmt.Errorf("error: %v", err)
}

// RedisExecuteWithValueError default redis response error message with value
func RedisExecuteWithValueError(value interface{}, err error) error {
	return fmt.Errorf("error with parameter %v: %v", value, err)
}

// RedisInfo default redis info message
func RedisInfo(msg string, values map[string]interface{}) string {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Sprintf("%v %v", msg, valueString)
}
