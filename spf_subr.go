package spf

import (
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"
)

// fetchTXT fetch all TXT records for domain d
func fetchTXT(ctx *Context, d string) ([]string, error) {
	debug("fetchtxt")
	if d == "" {
		return nil, fmt.Errorf("nil domain")
	}
	if ctx == nil {
		return nil, fmt.Errorf("nil context")
	}
	raw, err := ctx.r.LookupTXT(d)
	if err != nil {
		return nil, errors.Wrap(err, "lookuptxt")
	}
	return raw, nil
}

// fetchTXT fetch all TXT records for domain d
func fetchMX(ctx *Context, d string) ([]string, error) {
	var mxs []string

	debug("fetchmx")
	if d == "" {
		return nil, fmt.Errorf("nil domain")
	}
	if ctx == nil {
		return nil, fmt.Errorf("nil context")
	}
	raw, err := ctx.r.LookupMX(d)
	if err != nil {
		return nil, errors.Wrap(err, "lookupmx")
	}

	for _, mx := range raw {
		mxs = append(mxs, mx.Host)
	}
	return mxs, nil
}

// getSPF returns the first SPF record
func getSPF(rr []string) string {
	debug("getspf")
	for _, r := range rr {
		// We use only SPF (v=spf1), not SenderID (spf2.0/)
		if strings.HasPrefix(r, "v=spf1") {
			return r
		}
	}
	return ""
}

const (
	kwA = (1 << iota)
	kwMX
	kwPTR
	kwEXISTS
	kwINCLUDE
	kwIP4
	kwIP6
)

func parseTXT(txt string) ([]Entry, error) {
	// Split
	fields := strings.Fields(txt)
	if len(fields) == 0 || fields[0] != "v=spf1" {
		return nil, fmt.Errorf("wrong format %s", txt)
	}

	return nil, nil
}
