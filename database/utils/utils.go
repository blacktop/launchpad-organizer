package utils

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

var (
	normalPadding    = cli.Default.Padding
	doublePadding    = normalPadding * 2
	triplePadding    = normalPadding * 3
	quadruplePadding = normalPadding * 4
)

// Indent indents apex log line
func Indent(f func(s string)) func(string) {
	return func(s string) {
		cli.Default.Padding = doublePadding
		f(s)
		cli.Default.Padding = normalPadding
	}
}

// DoubleIndent double indents apex log line
func DoubleIndent(f func(s string)) func(string) {
	return func(s string) {
		cli.Default.Padding = triplePadding
		f(s)
		cli.Default.Padding = normalPadding
	}
}

// TripleIndent triple indents apex log line
func TripleIndent(f func(s string)) func(string) {
	return func(s string) {
		cli.Default.Padding = quadruplePadding
		f(s)
		cli.Default.Padding = normalPadding
	}
}

// StringInSlice finds string in array
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// AppendIfMissing with append a string to a array if it doesn't already contain that string
func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func checkError(err error) {
	if err != nil {
		log.WithError(err).Fatal("failed")
	}
}
