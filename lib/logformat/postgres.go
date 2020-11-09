package logformat

import "fmt"

// PostgresQueryResponseError default postgres query get response error message
func PostgresQueryResponseError(err error) error {
	return fmt.Errorf("postgres query response error: %v", err)
}

// PostgresScanError default postgres scan response error message
func PostgresScanError(err error) error {
	return fmt.Errorf("postgres scan error: %v", err)
}

// PostgresUnmarshalError default postgres unmarshal response error message
func PostgresUnmarshalError(err error) error {
	return fmt.Errorf("postgres unmarshal error: %v", err)
}
