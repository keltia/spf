# README.md

# Status

[![GitHub release](https://img.shields.io/github/release/keltia/spf.svg)](https://github.com/keltia/spf/releases)
[![GitHub issues](https://img.shields.io/github/issues/keltia/spf.svg)](https://github.com/keltia/spf/issues)
[![Go Version](https://img.shields.io/badge/go-1.10-blue.svg)](https://golang.org/dl/)
[![Build Status](https://travis-ci.org/keltia/spf.svg?branch=master)](https://travis-ci.org/keltia/spf)
[![GoDoc](http://godoc.org/github.com/keltia/spf?status.svg)](http://godoc.org/github.com/keltia/spf)
[![SemVer](http://img.shields.io/SemVer/2.0.0.png)](https://semver.org/spec/v2.0.0.html)
[![License](https://img.shields.io/pypi/l/Django.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![Go Report Card](https://goreportcard.com/badge/github.com/keltia/spf)](https://goreportcard.com/report/github.com/keltia/spf)

A few functions related to SPF usage.

# Installation

As with many Go utilities, a simple

    go get github.com/keltia/spf

is enough to fetch, build and install.

# Dependencies

* Go >= 1.10

Only standard Go modules are used.  I use Semantic Versioning for all my modules.

# Usage

SYNOPSIS
``` go

    You can have more verbose output and debug by using these functions:
    
    spf.SetVerbose()
    spf.SetDebug()      // implies verbose
    ...
    spf.Reset()         // both flags are cleared
```

# Tests

I'm trying to get to 100% coverage but some error cases are more difficult to create.

## License

This is released under the BSD 2-Clause license.  See `LICENSE.md`.

# Contributing

This project is an open Open Source project, please read `CONTRIBUTING.md`.

# Feedback

We welcome pull requests, bug fixes and issue reports.

Before proposing a large change, first please discuss your change by raising an issue.

I use Git Flow for this package so please use something similar or the usual github workflow.

1. Fork it ( https://github.com/keltia/spf/fork )
2. Checkout the develop branch (`git checkout develop`)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
