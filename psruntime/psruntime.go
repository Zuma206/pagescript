package psruntime

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type PSRuntime struct{}

func NewPSRuntime() *PSRuntime {
	return &PSRuntime{}
}

var passes = []NodeHandlers{
	evalNodeHandlers,
}

func (psr *PSRuntime) Run(input io.Reader, output io.Writer) error {
	document, err := html.Parse(input)
	if err != nil {
		return fmt.Errorf("failed to parse html: %w", err)
	}
	psc := &PSContext{
		output:    output,
		psRuntime: psr,
	}
	for _, passHandlers := range passes {
		psc.handlers = passHandlers
		if err := psc.RunNode(document); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintln(output); err != nil {
		return err
	}
	return nil
}
