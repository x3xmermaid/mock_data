package logformat

import (
	"fmt"
)

// InfluxConvertValueError default influx convert value error message
func InfluxConvertValueError(varName string, err error) error {
	return fmt.Errorf("influxdb unable to convert %v: %v", varName, err)
}

// InfluxGetResponseError default influx query get response error message
func InfluxGetResponseError(err error) error {
	return fmt.Errorf("influxdb GetResponse error: %v", err)
}

// InfluxQueryResponseError default influx query get response error message
func InfluxQueryResponseError(err error) error {
	return fmt.Errorf("influxdb query response error: %v", err)
}

// InfluxTimeIntervalError default influx time range interval error message
func InfluxTimeIntervalError(from, to int, err error) error {
	return fmt.Errorf("influxdb bad range time from %v - to %v: %v", from, to, err)
}

// InfluxQueryDataNotFound default influx query data not found message
func InfluxQueryDataNotFound() error {
	return fmt.Errorf("influxdb data not found")
}
