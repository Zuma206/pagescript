package psruntime

import (
	"fmt"

	"golang.org/x/net/html"
)

var evalNodeHandlers = &NodeHandlerRegistry{
	nodeHandlers: NodeHandlerMap{
		html.DocumentNode: evalChildren,
		html.DoctypeNode:  evalDoctype,
	},
}

func evalChildren(psc *PSContext, node *html.Node) error {
	for child := range node.ChildNodes() {
		if err := psc.RunNode(child); err != nil {
			return err
		}
	}
	return nil
}

func evalDoctype(psc *PSContext, node *html.Node) error {
	_, err := fmt.Fprintf(psc.Output(), "<!DOCTYPE %s>", node.Data)
	return err
}
