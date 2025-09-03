package psruntime

import "golang.org/x/net/html"

type NodeHandlerMap map[html.NodeType]NodeHandler

type NodeHandlerRegistry struct {
	nodeHandlers NodeHandlerMap
}

func NewNodeHandlers(handlers NodeHandlerMap) *NodeHandlerRegistry {
	return &NodeHandlerRegistry{handlers}
}

func (registry *NodeHandlerRegistry) SetNodeHandler(nodeType html.NodeType, handler NodeHandler) {
	registry.nodeHandlers[nodeType] = handler
}

func (registry *NodeHandlerRegistry) Get(nodeType html.NodeType) (NodeHandler, bool) {
	handler, ok := registry.nodeHandlers[nodeType]
	return handler, ok
}
