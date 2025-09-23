package stdlib

import (
	"github.com/Zuma206/pagescript/psruntime"
	"golang.org/x/net/html"
)

func OpenElements(psr *psruntime.PSRuntime) error {
	elements := psr.Engine().NewObject()
	psr.ScanHandlers().SetAttrHandler("id", func(_ *psruntime.PSContext, node *html.Node, id string) error {
		element := psr.Engine().NewObject()
		if err := element.Set("type", node.Data); err != nil {
			return err
		}
		return elements.Set(id, element)
	})
	return psr.Engine().Set("elements", elements)
}
