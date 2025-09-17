package psruntime

import (
	"fmt"
	"io"

	"github.com/Zuma206/pagescript/eventloop"
	"github.com/dop251/goja"
	"golang.org/x/net/html"
)

type PSRuntime struct {
	eventloop *eventloop.Eventloop
	engine    *goja.Runtime
	log       io.Writer
}

func (runtime *PSRuntime) Eventloop() *eventloop.Eventloop {
	return runtime.eventloop
}

func (runtime *PSRuntime) Engine() *goja.Runtime {
	return runtime.engine
}

func (runtime *PSRuntime) Log() io.Writer {
	return runtime.log
}

var passes = []NodeHandlers{
	evalNodeHandlers,
}

func (runtime *PSRuntime) Run(input io.Reader, output io.Writer) error {
	document, err := html.Parse(input)
	if err != nil {
		return fmt.Errorf("failed to parse html: %w", err)
	}
	psc := &PSContext{
		output:  output,
		runtime: runtime,
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
