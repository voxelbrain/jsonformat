package main

import (
	"io"
	"text/template"
)

type Formatter interface {
	Execute(w io.Writer, data interface{}) error
}

type FormatterFunc func(w io.Writer, data interface{}) error

func (ff FormatterFunc) Execute(w io.Writer, data interface{}) error {
	return ff(w, data)
}

type Compiler func(format string) (Formatter, error)
type Format struct {
	Compiler
	Description string
}

var Formats = map[string]Format{
	"template": Format{
		Compiler: func(format string) (Formatter, error) {
			return template.New("jsonformat").Parse(format)
		},
		Description: `Thin wrapper around Go's templating language. (http://golang.org/pkg/text/template/)`,
	},
}
