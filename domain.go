package spf

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	myVersion = "0.0.1"
)

// Context is used to switch resolvers (and test/mock)
type Context struct {
	r Resolver
}

// Domain represents a given SPF RR and the IPs behind
type Domain struct {
	ctx  *Context
	Name string
	Raw  string
}

// NewDomain creates a Domain object
func NewDomain(dom string) (*Domain, error) {
	debug("newdomain=%s", dom)
	if dom == "" {
		return &Domain{}, fmt.Errorf("empty domain")
	}

	ctx := &Context{RealResolver{}}
	return &Domain{ctx: ctx, Name: dom}, nil
}

// Fetch gets the SPF TXT RR
func (d *Domain) Fetch() error {
	raw, err := fetchTXT(d.ctx, d.Name)
	if err != nil {
		return errors.Wrap(err, "fetchtxt")
	}
	debug("all txt=%v", raw)
	rawspf := getSPF(raw)
	if rawspf == "" {
		return fmt.Errorf("no spf")
	}

	d.Raw = rawspf
	return nil
}

// FetchMX gets the MX RR
func (d *Domain) FetchMX() ([]string, error) {
	raw, err := fetchMX(d.ctx, d.Name)
	if err != nil {
		return []string{}, errors.Wrap(err, "fetchmx")
	}
	debug("all txt=%v", raw)

	return raw, nil
}

/*
SPF entries

v=spf1 entry* <action>

entry are

a
exists:
include:
ip4:
ip6:
mx
ptr
redirect:

XXX macros are not supported
*/

// Unroll does recursively fetch & unroll SPF (by default
func (d *Domain) Unroll(limit int) (*Result, error) {
	// If not yet fetched
	if d.Raw == "" {
		err := d.Fetch()
		if err != nil {
			return nil, errors.Wrap(err, "unroll/fetch")
		}
	}

	r := NewResult()
	err := r.parseTXT(d.Name)

	return r, err
}

var (
	fVerbose = false
	fDebug   = false
)

// SetVerbose sets the mode
func SetVerbose() {
	fVerbose = true
}

// SetDebug sets the mode too
func SetDebug() {
	fDebug = true
	fVerbose = true
}

// Reset is for the two flags
func Reset() {
	fDebug = false
	fVerbose = false
}

// Version returns our current version
func Version() string {
	return myVersion
}
