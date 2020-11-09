package inputvalidator

import (
	"net"
	"regexp"
)

const (
	//HostnameRegex hostname regex format
	HostnameRegex = `^(([a-zA-Z]|[a-zA-Z][a-zA-Z0-9-]*[a-zA-Z0-9])\.)*([A-Za-z]|[A-Za-z][A-Za-z0-9-]*[A-Za-z0-9])$`
)

// IsIPHostname validates ipv4, ipv6 and hostname string format
func IsIPHostname(input string) bool {
	validIP := net.ParseIP(input)
	if validIP == nil {
		matched, err := regexp.MatchString(HostnameRegex, input)
		if err != nil {
			return false
		}

		return matched
	}

	return true
}
