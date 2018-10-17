package spf

import (
	"fmt"
	"net"
)

// Resolver is the main interface we use
type Resolver interface {
	LookupTXT(addr string) ([]string, error)
}

// NullResolver is empty
type NullResolver struct{}

// LookupTXT always return a good and fixed answer
func (NullResolver) LookupTXT(addr string) ([]string, error) {
	return []string{addr}, nil
}

// RealResolver will call the real one
type RealResolver struct{}

// LookupTXT use the real "net" function
func (r RealResolver) LookupTXT(addr string) ([]string, error) {
	return net.LookupTXT(addr)
}

// ErrorResolver always returns an error
type ErrorResolver struct{}

// LookupTXT is for testing errors
func (ErrorResolver) LookupTXT(s string) ([]string, error) {
	return []string{}, fmt.Errorf("error")
}
