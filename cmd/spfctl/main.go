// main.go

/*
Small demonstration program for github.com/keltia/spf
*/

package main

import (
	"fmt"

	"github.com/keltia/spf"
)

const (
	myName = "spfctl"
)

func main() {
	fmt.Printf("%s API version/%s\n", myName, spf.Version())
}
