package spf

import (
	"net"
)

// Resolver is the main interface we use
type Resolver interface {
	LookupTXT(addr string) ([]string, error)
}

// NullResolver is empty
type NullResolver struct{}

// LookupAddr always return a good and fixed answer
func (NullResolver) LookupTXT(addr string) ([]string, error) {
	return []string{addr}, nil
}

// RealResolver will call the real one
type RealResolver struct{}

// LookupAddr use the real "net" function
func (r RealResolver) LookupTXT(addr string) ([]string, error) {
	return net.LookupTXT(addr)
}
