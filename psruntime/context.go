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
	handlers NodeHandlers
	runtime  *PSRuntime
	output   io.Writer
}

func (ctx *PSContext) Output() io.Writer {
	return ctx.output
}

func (ctx *PSContext) RunNode(node *html.Node) error {
	handler, ok := ctx.handlers.Get(node.Type)
	if !ok {
		return fmt.Errorf("can't handle node of type %d", node.Type)
	}
	return handler(ctx, node)
}
