package psruntime

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type NodeHandler func(psc *PSContext, node *html.Node) error
type NodeHandlers interface {
	Get(nodeType html.NodeType) (NodeHandler, bool)
}

type PSContext struct {
	handlers  NodeHandlers
	psRuntime *PSRuntime
	output    io.Writer
}

func (psc *PSContext) Output() io.Writer {
	return psc.output
}

func (psc *PSContext) RunNode(node *html.Node) error {
	handler, ok := psc.handlers.Get(node.Type)
	if !ok {
		return fmt.Errorf("can't handle node of type %d", node.Type)
	}
	return handler(psc, node)
}
