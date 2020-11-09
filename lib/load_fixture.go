package lib

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Loadfixture loads a fixture file and return to byte array
func Loadfixture(f string) []byte {
	pwd, _ := os.Getwd()
	p := filepath.Join(pwd, ".", "_fixtures", f)
	c, _ := ioutil.ReadFile(p)
	return c
}

// LoadMockData loads a fixture file and return to byte array
func LoadMockData(f string) []byte {
	pwd, _ := os.Getwd()
	p := filepath.Join(pwd, ".", "mock-data", f)
	c, _ := ioutil.ReadFile(p)
	return c
}
