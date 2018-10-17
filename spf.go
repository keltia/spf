package spf

import (
	"fmt"
	"net"
	"strings"

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
	IPs  Blocks
}

// Blocks is the final list of all IPs
type Blocks []net.IPNet

// Entry is used when parsing SPF records
type Entry struct {
	t string
	v string
}

// NewDomain creates a Domain object
func NewDomain(dom string) (*Domain, error) {
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
	rawspf := getSPF(raw)
	if rawspf == "" {
		return fmt.Errorf("no spf")
	}

	d.Raw = rawspf
	return nil
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

XXX macros are not supported
*/

// Recursively fetch & unroll SPF (by default
func (d *Domain) Unroll(limit int) (Blocks, error) {
	// If not yet fetched
	if d.Raw == "" {
		err := d.Fetch()
		if err != nil {
			return nil, errors.Wrap(err, "unroll/fetch")
		}
	}

	// Split
	fields := strings.Fields(d.Raw)
	if len(fields) == 0 || fields[0] != "v=spf1" {
		return nil, fmt.Errorf("wrong format %s", d.Raw)
	}

	return nil, nil
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
