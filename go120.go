// Package version_test
//go:build go1.20

package version_test

import "github.com/klauspost/compress/zstd"

type i struct {
	zstd.Decoder
}

func (i) PrintVersion() string {
	return "1.20"
}

func NewPrintVersion() VersionPrinter {
	return i{}
}
