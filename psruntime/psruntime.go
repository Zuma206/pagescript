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

type PSContext struct {
	handlers  NodeHandlers
	psRuntime *PSRuntime
	output    io.Writer
}

func (psc *PSContext) Output() io.Writer {
	return psc.output
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

type NodeHandler func(psc *PSContext, node *html.Node) error
type NodeHandlers interface {
	Get(nodeType html.NodeType) (NodeHandler, bool)
}

func (psc *PSContext) RunNode(node *html.Node) error {
	handler, ok := psc.handlers.Get(node.Type)
	if !ok {
		return fmt.Errorf("can't handle node of type %d", node.Type)
	}
	return handler(psc, node)
}
