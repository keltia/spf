package spf

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/yl2chen/cidranger"
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

// Result is the main object used to recurse and store
type Result struct {
	IPs []net.IPNet

	rec int
	dns int
	rn  cidranger.Ranger
}

func NewResult() *Result {
	return &Result{rn: cidranger.NewPCTrieRanger()}
}

// AppendMX gets all IP for the MXs
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

			present, _ := r.rn.Contains(net.ParseIP(aip))
			if !present {
				r.IPs = append(r.IPs, *ipb)
			} else {
				debug("included")
			}
		}
	}
	return nil
}

func readIP(m []string) *net.IPNet {
	net0 := m[3]
	if !strings.Contains(net0, "/") {
		net0 = net0 + netm[m[1]]
	}
	ip, ipb, _ := net.ParseCIDR(net0)
	debug("ipnet: %s-%s", ip, ipb)
	return ipb
}

// ParseTXT does the main SPF parsing
func (r *Result) parseTXT(dom string) error {
	if r.rec >= 10 {
		return fmt.Errorf("recursion limit")
	}

	d, _ := NewDomain(dom)
	d.Fetch()

	// Sanity control
	if r.rn == nil {
		r.rn = cidranger.NewPCTrieRanger()
	}

	txt := strings.Fields(d.Raw)

	// First checks
	if len(txt) == 0 || txt[0] != "v=spf1" {
		return fmt.Errorf("wrong format %s", txt)
	}

	// We are not parsing completely, just what interest us
	for _, f := range txt[1:] {
		if m := reIPS.FindStringSubmatch(f); m != nil {
			//
			debug("ip46: %s", m[3])

			present, _ := r.rn.Contains(net.ParseIP(m[3]))
			if !present {
				ipb := readIP(m)
				r.IPs = append(r.IPs, *ipb)
			} else {
				debug("included")
			}
		} else if m := reINC.FindStringSubmatch(f); m != nil {
			//
			debug("include: %s", m)

			r.rec++
			r.dns++
			// Recurse with include
			err := r.parseTXT(m[2])
			debug("error was %v", err)
		} else if reMXS.MatchString(f) {
			//
			debug("mx")

			r.AppendMX(d.Name)
		} else if reRDR.FindStringSubmatch(f); m != nil {
			//
			debug("redirect: %s", m)

			// If redirect= and *all, ignore
			if strings.HasSuffix(txt[len(txt)], "all") {
				continue
			}
			r.rec++
			r.dns++

			// Recurse with redirect
			err := r.parseTXT(m[1])
			debug("error was %v", err)
		} else {
			debug("nothing")
		}
	}

	return nil
}

func (r *Result) Condense() []cidranger.RangerEntry {
	rn := cidranger.NewPCTrieRanger()
	for _, ipn := range r.IPs {
		rn.Insert(cidranger.NewBasicRangerEntry(ipn))
	}
	res, _ := rn.ContainingNetworks(net.ParseIP("0.0.0.0/32"))
	debug("res=%v", res)
	return res
}

// String() is the basic Stringer
func (r *Result) String() string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("Max recursion level: %d\n", r.rec))
	str.WriteString(fmt.Sprintf("Total DNS calls: %d\n", r.dns))
	str.WriteString("IP Blocks:\n")
	for _, b := range r.IPs {
		str.WriteString(fmt.Sprintf("%s ", b.String()))
	}
	return str.String()
}
