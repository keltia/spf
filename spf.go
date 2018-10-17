package spf

import (
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"
)

type Context struct {
	r Resolver
}

type Domain struct {
	ctx  *Context
	Name string
	Raw  string
	IPs  Blocks
}

type Blocks []net.IPNet

type Entry struct {
	t string
	v string
}

func NewDomain(dom string) (*Domain, error) {
	if dom == "" {
		return &Domain{}, fmt.Errorf("empty domain")
	}

	ctx := &Context{RealResolver{}}
	return &Domain{ctx: ctx, Name: dom}, nil
}

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
