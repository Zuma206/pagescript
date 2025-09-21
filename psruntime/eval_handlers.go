package psruntime

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func newEvalHandlers() *NodeHandlerRegistry {
	return NewNodeHandlers(
		NodeHandlerMap{
			html.DocumentNode: handleChildren,
			html.DoctypeNode:  evalDoctype,
			html.ElementNode:  evalElement,
			html.TextNode:     evalText,
			html.CommentNode:  handleSkip,
		},
		ElementHandlerMap{
			"script": evalScript,
		},
		AttrHandlerMap{},
	)
}

func evalDoctype(ctx *PSContext, node *html.Node) error {
	_, err := fmt.Fprintf(ctx.Output(), "<!DOCTYPE %s>", node.Data)
	return err
}

func evalElement(ctx *PSContext, node *html.Node) error {
	if _, err := fmt.Fprintf(ctx.Output(), "<%s", node.Data); err != nil {
		return err
	}
	if err := evalAttrs(ctx, node); err != nil {
		return err
	}
	if voidElements.Has(node.Data) {
		if _, err := fmt.Fprint(ctx.Output(), "/>"); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprint(ctx.Output(), ">"); err != nil {
			return err
		}
		if err := handleChildren(ctx, node); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(ctx.Output(), "</%s>", node.Data); err != nil {
			return err
		}
	}
	return nil
}

func evalAttrs(ctx *PSContext, node *html.Node) error {
	for _, attr := range node.Attr {
		if _, err := fmt.Fprintf(ctx.Output(), " %s=\"%s\"", attr.Key, attr.Val); err != nil {
			return err
		}
	}
	return nil
}

func evalText(ctx *PSContext, node *html.Node) error {
	_, err := fmt.Fprint(ctx.Output(), strings.TrimSpace(node.Data))
	return err
}

func evalScript(ctx *PSContext, node *html.Node) error {
	return ctx.Runtime().Eventloop().Block(func() error {
		_, err := ctx.Runtime().Engine().RunString(node.FirstChild.Data)
		return err
	})
}
