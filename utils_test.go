package spf

import (
	"testing"
)

func TestVerbose_No(t *testing.T) {
	verbose("no")
}

func TestVerbose_Yes(t *testing.T) {
	SetVerbose()
	verbose("yes")
	fVerbose = false
}

func TestDebug_No(t *testing.T) {
	debug("no")
}

func TestDebug_Yes(t *testing.T) {
	SetDebug()
	debug("yes")
	fDebug = false
}
