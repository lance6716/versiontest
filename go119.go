// Package version_test
//go:build !go1.20

package version_test

import "github.com/DataDog/zstd"

type i struct {
	zstd.Writer
}

func (i) PrintVersion() string {
	return "not 1.20"
}

func NewPrintVersion() VersionPrinter {
	return i{}
}
