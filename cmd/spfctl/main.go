// main.go

/*
Small demonstration program for github.com/keltia/spf
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/keltia/spf"
	"github.com/pkg/errors"
)

const (
	myName = "spfctl"
)

var (
	fDebug   bool
	fVerbose bool
)

func init() {
	flag.BoolVar(&fDebug, "D", false, "Debug mode")
	flag.BoolVar(&fVerbose, "v", false, "Verbose mode")
}

var ErrUsage = errors.New("must exit")

func Setup() error {
	fmt.Printf("%s API version/%s\n", myName, spf.Version())
	if fVerbose {
		spf.SetVerbose()
	}

	if fDebug {
		spf.SetDebug()
	}

	if len(flag.Args()) < 1 {
		log.Println("You must supply at least one domain!")
		return ErrUsage
	}
	return nil
}

func main() {
	flag.Parse()

	err := Setup()
	if err == ErrUsage {
		os.Exit(1)
	}

	str := flag.Arg(0)
	d, err := spf.NewDomain(str)
	if err != nil {
		log.Fatalf("Error creating Domain: %v")
	}

	err = d.Fetch()
	fmt.Printf("SPF field:\n%s\n", d.Raw)

	res, err := d.Unroll(0)
	fmt.Printf("IP blocks:\n%s\n", res)
}
