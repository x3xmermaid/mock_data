package logformat

import (
	"encoding/json"
	"fmt"
)

const (
	// InternalServerError default message
	InternalServerError = "internal server error"
	// BadRequest default message
	BadRequest = "bad request"
	// NotFound default message
	NotFound = "not found"
)

// HandlerInfluxQueryError default influx query error message
func HandlerInfluxQueryError(msg string, values interface{}, err error) error {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Errorf("Unable to %v %v: %v [Influx]", msg, valueString, err)
}

// HandlerInfluxQuerySuccess default influx query error message
func HandlerInfluxQuerySuccess(msg string, values interface{}) string {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Sprintf("Success %v %v [Influx]", msg, valueString)
}

// HandlerPostgresQueryError default postgres query error message
func HandlerPostgresQueryError(msg string, values interface{}, err error) error {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Errorf("Unable to %v %v: %v [Postgres]", msg, valueString, err)
}

// HandlerPostgresQuerySuccess default postgres query error message
func HandlerPostgresQuerySuccess(msg string, values interface{}) string {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Sprintf("Success %v %v [Postgres]", msg, valueString)
}

// HandlerMongoQueryError default mongo query error message
func HandlerMongoQueryError(msg string, values interface{}, err error) error {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Errorf("Unable to %v %v: %v [Mongo]", msg, valueString, err)
}

// HandlerMongoQuerySuccess default mongo query error message
func HandlerMongoQuerySuccess(msg string, values interface{}) string {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Sprintf("Success %v %v [Mongo]", msg, valueString)
}

// HandlerRedisExecuteError default redis query error message
func HandlerRedisExecuteError(msg string, values interface{}, err error) error {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Errorf("Unable to %v %v: %v [Redis]", msg, valueString, err)
}

// HandlerRedisExecuteSuccess default redis query success message
func HandlerRedisExecuteSuccess(msg string, values interface{}) string {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Sprintf("Success %v %v [Redis]", msg, valueString)
}

// HandlerConvertValueError default handler convert value error message
func HandlerConvertValueError(key string, value interface{}, err error) error {
	return fmt.Errorf("Unable to convert %v value %v: %v", key, value, err)
}

// HandlerValidateIPHostname default ip/hostname validator error message
func HandlerValidateIPHostname(value interface{}) error {
	return fmt.Errorf("Invalid ip/hostname value %v", value)
}

// HandlerInvalidParameter default invalid parameter error message
func HandlerInvalidParameter(key string) error {
	return fmt.Errorf("Invalid %v parameter", key)
}

// HandlerInvalidRequest default request body decode error message
func HandlerInvalidRequest(err error) error {
	return fmt.Errorf("Invalid request: %v", err)
}

// HandlerCheckDataNotExist default request check data not exist error message
func HandlerCheckDataNotExist(key, value string) error {
	return fmt.Errorf("Data %v value %v not exist", key, value)
}

// HandlerObjectInitializationError default object initialization error message
func HandlerObjectInitializationError(values string, err error) error {
	return fmt.Errorf("Unable to initiate object %v: %v", values, err)
}

// HanderNotFound default data not found error message
func HanderNotFound(key, value string) error {
	return fmt.Errorf("Data %v value %v not found", key, value)
}

// HandlerEmptyValue default empty parameter error message
func HandlerEmptyValue(funcname, field string) error {
	return fmt.Errorf("%v invalid parameter: %v is empty", funcname, field)
}

// HandlerInputReachLimit default reach limitation error message
func HandlerInputReachLimit() error {
	return fmt.Errorf("Cannot insert, has already reached quota limitation")
}

// HandlerErrorIOCopy default cannot copy error message
func HandlerErrorIOCopy(err error) error {
	return fmt.Errorf("Cannot copy bytes data: %v", err)
}

// HandlerNodeInterfaceNotExist default node interface not exist message
func HandlerNodeInterfaceNotExist(key, value interface{}) error {
	return fmt.Errorf("Node %v with interface %v not found", key, value)
}

// HandlerGatewayAlreadyExist default gateway already exist message
func HandlerGatewayAlreadyExist(key, value interface{}) error {
	return fmt.Errorf("Node %v with gateway %v already exist", key, value)
}

// HandlerNodeAlreadyExist default node already exist message
func HandlerNodeAlreadyExist(key interface{}) error {
	return fmt.Errorf("Node with ip address %v already exist", key)
}

// HandlerProbeAlreadyExist default node already exist message
func HandlerProbeAlreadyExist(key interface{}) error {
	return fmt.Errorf("Probe %v already exist", key)
}
