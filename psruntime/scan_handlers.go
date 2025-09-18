package psruntime

import "golang.org/x/net/html"

func newScanHandlers() *NodeHandlerRegistry {
	return NewNodeHandlers(
		NodeHandlerMap{
			html.DocumentNode: handleChildren,
			html.DoctypeNode:  handleSkip,
			html.CommentNode:  handleSkip,
		},
		ElementHandlerMap{},
	)
}
