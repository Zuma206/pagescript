package psruntime

func newScanHandlers() *NodeHandlerRegistry {
	return NewNodeHandlers(
		NodeHandlerMap{},
		ElementHandlerMap{},
	)
}

