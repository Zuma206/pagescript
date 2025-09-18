package psruntime

import "golang.org/x/net/html"

func handleChildren(ctx *PSContext, node *html.Node) error {
	for child := range node.ChildNodes() {
		if err := ctx.RunNode(child); err != nil {
			return err
		}
	}
	return nil
}

func handleSkip(_ *PSContext, _ *html.Node) error {
	return nil
}
