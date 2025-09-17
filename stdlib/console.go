package stdlib

import (
	"fmt"

	"github.com/Zuma206/pagescript/psruntime"
	"github.com/dop251/goja"
)

func OpenConsole(psr *psruntime.PSRuntime) error {
	console := psr.Engine().NewObject()
	console.Set("log", consoleLog(psr))
	return psr.Engine().Set("console", console)
}

func consoleLog(psr *psruntime.PSRuntime) any {
	return func(values ...goja.Value) {
		for i, value := range values {
			if i != 0 {
				fmt.Fprint(psr.Log(), " ")
			}
			fmt.Fprint(psr.Log(), value.String())
		}
		fmt.Fprintln(psr.Log())
	}
}
