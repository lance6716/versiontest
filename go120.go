// Package version_test
//go:build go1.20

package version_test

type i struct{}

func (i) PrintVersion() string {
	return "1.20"
}

func NewPrintVersion() VersionPrinter {
	return i{}
}
