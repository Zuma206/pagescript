package psruntime

import "golang.org/x/net/html"

func newScanHandlers() *NodeHandlerRegistry {
	return NewNodeHandlers(
		NodeHandlerMap{
			html.DocumentNode: handleChildren,
			html.ElementNode:  handleChildren,
			html.DoctypeNode:  handleSkip,
			html.CommentNode:  handleSkip,
			html.TextNode:     handleSkip,
		},
		ElementHandlerMap{},
		AttrHandlerMap{},
	)
}
