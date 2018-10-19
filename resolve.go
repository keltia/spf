package spf

import (
	"fmt"
	"net"
)

// Resolver is the main interface we use
type Resolver interface {
	LookupHost(host string) ([]string, error)
	LookupMX(addr string) ([]*net.MX, error)
	LookupTXT(addr string) ([]string, error)
}

// NullResolver is empty
type NullResolver struct{}

// LookupTXT always return a good and fixed answer
func (NullResolver) LookupTXT(addr string) ([]string, error) {
	return []string{addr}, nil
}

// LookupHost is for forward resolving
func (NullResolver) LookupHost(host string) ([]string, error) {
	return []string{}, nil
}

// LookupMX use the real "net" function
func (NullResolver) LookupMX(addr string) ([]*net.MX, error) {
	return []*net.MX{}, nil
}

// RealResolver will call the real one
type RealResolver struct{}

// LookupTXT use the real "net" function
func (RealResolver) LookupTXT(addr string) ([]string, error) {
	return net.LookupTXT(addr)
}

// LookupMX use the real "net" function
func (RealResolver) LookupMX(addr string) ([]*net.MX, error) {
	return net.LookupMX(addr)
}

// LookupHost is for forward resolving
func (RealResolver) LookupHost(host string) ([]string, error) {
	return net.LookupHost(host)
}

// ErrorResolver always returns an error
type ErrorResolver struct{}

// LookupTXT is for testing errors
func (ErrorResolver) LookupTXT(s string) ([]string, error) {
	return []string{}, fmt.Errorf("error")
}

// LookupHost is for forward resolving
func (ErrorResolver) LookupHost(host string) ([]string, error) {
	return []string{}, fmt.Errorf("error")
}

// LookupMX use the real "net" function
func (r ErrorResolver) LookupMX(addr string) ([]*net.MX, error) {
	return net.LookupMX(addr)
}
