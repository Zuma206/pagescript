package psruntime

import "golang.org/x/net/html"

type NodeHandlerMap map[html.NodeType]NodeHandler
type ElementHandlerMap map[string]NodeHandler
type NodeValueHandler func(psc *PSContext, node *html.Node, value string) error
type AttrHandlerMap map[string]NodeValueHandler

type NodeHandlerRegistry struct {
	nodeHandlers    NodeHandlerMap
	elementHandlers ElementHandlerMap
	attrHandlers    AttrHandlerMap
}

func NewNodeHandlers(
	nodeHandlers NodeHandlerMap, elementHandlers ElementHandlerMap, attrHandlers AttrHandlerMap) *NodeHandlerRegistry {
	return &NodeHandlerRegistry{nodeHandlers, elementHandlers, attrHandlers}
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
	if err := registry.handleAttrs(psc, node); err != nil {
		return err
	}
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

func (registry *NodeHandlerRegistry) handleAttrs(psc *PSContext, node *html.Node) error {
	for _, attr := range node.Attr {
		if handler, ok := registry.attrHandlers[attr.Key]; ok {
			if err := handler(psc, node, attr.Val); err != nil {
				return err
			}
		}
	}
	return nil
}

func (registry *NodeHandlerRegistry) SetElementHandler(elementType string, handler NodeHandler) {
	registry.elementHandlers[elementType] = handler
}

func (registry *NodeHandlerRegistry) SetAttrHandler(attrKey string, handler NodeValueHandler) {
	registry.attrHandlers[attrKey] = handler
}
