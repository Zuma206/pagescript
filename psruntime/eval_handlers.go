package psruntime

import (
	"fmt"

	"golang.org/x/net/html"
)

var evalNodeHandlers = &NodeHandlerRegistry{
	nodeHandlers: NodeHandlerMap{
		html.DocumentNode: evalChildren,
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
