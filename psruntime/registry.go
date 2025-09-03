package psruntime

import "golang.org/x/net/html"

type NodeHandlerMap map[html.NodeType]NodeHandler
type ElementHandlerMap map[string]NodeHandler

type NodeHandlerRegistry struct {
	nodeHandlers    NodeHandlerMap
	elementHandlers ElementHandlerMap
}

func NewNodeHandlers(nodeHandlers NodeHandlerMap, elementHandlers ElementHandlerMap) *NodeHandlerRegistry {
	return &NodeHandlerRegistry{nodeHandlers, elementHandlers}
}

func (registry *NodeHandlerRegistry) SetNodeHandler(nodeType html.NodeType, handler NodeHandler) {
	registry.nodeHandlers[nodeType] = handler
}

func (registry *NodeHandlerRegistry) Get(nodeType html.NodeType) (NodeHandler, bool) {
	if nodeType == html.ElementNode {
		return registry.handleElement, true
	}
	handler, ok := registry.nodeHandlers[nodeType]
	return handler, ok
}

func (registry *NodeHandlerRegistry) handleElement(psc *PSContext, node *html.Node) error {
	handler := registry.elementHandlers[node.Data]
	if handler != nil {
		return handler(psc, node)
	}
	defaultHandler := registry.nodeHandlers[html.ElementNode]
	if defaultHandler != nil {
		return defaultHandler(psc, node)
	}
	return nil
}

func (registry *NodeHandlerRegistry) SetElementHandler(elementType string, handler NodeHandler) {
	registry.elementHandlers[elementType] = handler
}
