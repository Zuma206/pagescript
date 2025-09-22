package stdlib

import (
	"github.com/Zuma206/pagescript/datatypes"
	"github.com/dop251/goja"
)

var nativeBindings = datatypes.NewWeakMap[goja.Object, any]()

func Native[T any](this *goja.Object) T {
	native, _ := nativeBindings.Get(this)
	return native.(T)
}

func Bind(this *goja.Object, native any) {
	nativeBindings.Set(this, native)
}
