package inputvalidator_test

import (
	ninputvalidator "ketitik/netmonk/mock-app-data/lib/inputvalidator"
	"testing"
)

func TestIsIPHostname(t *testing.T) {
	t.Run("IPV4 NOK", func(t *testing.T) {
		input := "127.0.0.1.1"
		matched := ninputvalidator.IsIPHostname(input)
		if matched {
			t.Errorf("It should not be match")
		}
	})

	t.Run("IPV6 NOK", func(t *testing.T) {
		input := "2001:0db8:85a3:0000:0000:8a2e:0370:7334:1111"
		matched := ninputvalidator.IsIPHostname(input)
		if matched {
			t.Errorf("It should not be match")
		}
	})

	t.Run("Hostname NOK", func(t *testing.T) {
		input := "#.netmonk.id"
		matched := ninputvalidator.IsIPHostname(input)
		if matched {
			t.Errorf("It should not be match")
		}
	})

	t.Run("IPV4 OK", func(t *testing.T) {
		input := "127.0.0.1"
		matched := ninputvalidator.IsIPHostname(input)
		if !matched {
			t.Errorf("It should be match")
		}
	})

	t.Run("IPV6 OK", func(t *testing.T) {
		input := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
		matched := ninputvalidator.IsIPHostname(input)
		if !matched {
			t.Errorf("It should be match")
		}
	})

	t.Run("Hostname OK", func(t *testing.T) {
		input := "router1.netmonk.id"
		matched := ninputvalidator.IsIPHostname(input)
		if !matched {
			t.Errorf("It should be match")
		}
	})
}
