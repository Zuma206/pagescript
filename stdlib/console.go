package stdlib

import (
	"github.com/Zuma206/pagescript/psruntime"
	"github.com/dop251/goja"
)

func OpenConsole(psr *psruntime.PSRuntime) error {
	console := psr.Engine().NewObject()
	console.Set("log", ConsoleLog(psr))
	return psr.Engine().Set("console", console)
}

func ConsoleLog(psr *psruntime.PSRuntime) any {
	return func(values ...goja.Value) {
		for i, value := range values {
			if i != 0 {
				print(" ")
			}
			print(value.String())
		}
		println()
	}
}
