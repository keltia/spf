package spf

import (
	"fmt"
	"net"
    "regexp"
    "strings"
)

const (
    matchQ   = `^([+\-\~\?])*`
    matchIP  = `^ip(4|6)([:/])(.*)$`
    matchINC = `^include([:/])(.*)$`
    matchMX  = `^mx$`
    matchRD  = `^redirect=(.*)$`
)

var (
    netm = map[string]string{"4": "/32", "6": "/128"}

    reIPS = regexp.MustCompile(matchIP)
    reINC = regexp.MustCompile(matchINC)
    reMXS = regexp.MustCompile(matchMX)
    reRDR = regexp.MustCompile(matchRD)
)

type Result struct {
	IPs Blocks

	rec int
	dns int
}

func (r *Result) AppendMX(dom string) error {
	d, _ := NewDomain(dom)
	mx, _ := d.FetchMX()
	r.dns += len(mx) + 1

	debug("mx=%v", mx)
	for _, m := range mx {
		alist, _ := d.ctx.r.LookupHost(m)
		debug("alist=%v", alist)
		for _, aip := range alist {
			// Trick to get proper string whether it is v4 or v6
			rip := net.ParseIP(aip)
			if rip.DefaultMask() == nil {
				aip = aip + "/128"
			} else {
				aip = aip + "/32"
			}
			_, ipb, err := net.ParseCIDR(aip)
			if err != nil {
				debug("%v for %v", err, aip)
			}
			r.IPs = append(r.IPs, *ipb)
		}
	}
	return nil
}

func (r *Result) parseTXT(dom string) (Blocks, error) {
	if r.rec >= 10 {
		return nil, fmt.Errorf("recursion limit")
	}

	d, _ := NewDomain(dom)
	d.Fetch()

	txt := strings.Fields(d.Raw)

	// First checks
	if len(txt) == 0 || txt[0] != "v=spf1" {
		return nil, fmt.Errorf("wrong format %s", txt)
	}

	// We are not parsing completely, just what interest us
	for _, f := range txt[1:] {
		if m := reIPS.FindStringSubmatch(f); m != nil {
			//
			debug("ip46: %s", m)

			net0 := m[3]
			if !strings.Contains(net0, "/") {
				net0 = net0 + netm[m[1]]
			}
			ip, ipb, err := net.ParseCIDR(net0)
			if err != nil {
				continue
			}
			debug("ipnet: %s-%s", ip, ipb)
			r.IPs = append(r.IPs, *ipb)
		} else if m := reINC.FindStringSubmatch(f); m != nil {
			//
			debug("include: %s", m)

			r.rec++
			r.dns++
			// Recurse with include
			b, err := r.parseTXT(m[2])
			if err != nil {
				continue
			}
			r.IPs = append(r.IPs, b...)
		} else if reMXS.MatchString(f) {
			//
			debug("mx")

			r.AppendMX(d.Name)
		} else if reRDR.MatchString(f) {
			//
			debug("redirect: %s", m)

			// If redirect= and *all, ignore
			if strings.HasSuffix(txt[len(txt)], "all") {
				continue
			}
		} else {
			debug("nothing")
		}
	}

	return r.IPs, nil
}