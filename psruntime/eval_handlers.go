package psruntime

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var evalNodeHandlers = NewNodeHandlers(
	NodeHandlerMap{
		html.DocumentNode: evalChildren,
		html.DoctypeNode:  evalDoctype,
		html.ElementNode:  evalElement,
		html.TextNode:     evalText,
	},
	ElementHandlerMap{},
)

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

func evalElement(psc *PSContext, node *html.Node) error {
	if _, err := fmt.Fprintf(psc.Output(), "<%s", node.Data); err != nil {
		return err
	}
	if err := evalAttrs(psc, node); err != nil {
		return err
	}
	if _, err := fmt.Fprint(psc.Output(), ">"); err != nil {
		return err
	}
	if err := evalChildren(psc, node); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(psc.Output(), "</%s>", node.Data); err != nil {
		return err
	}
	return nil
}

func evalAttrs(psc *PSContext, node *html.Node) error {
	for _, attr := range node.Attr {
		if _, err := fmt.Fprintf(psc.Output(), " %s=\"%s\"", attr.Key, attr.Val); err != nil {
			return err
		}
	}
	return nil
}

func evalText(psc *PSContext, node *html.Node) error {
	_, err := fmt.Fprint(psc.Output(), strings.TrimSpace(node.Data))
	return err
}
