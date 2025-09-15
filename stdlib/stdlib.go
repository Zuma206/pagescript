package stdlib

import (
	"github.com/Zuma206/pagescript/psruntime"
)

func Open(psr *psruntime.PSRuntime) error {
	return openAll(psr, []OpenFunction{
		OpenConsole,
	})
}

type OpenFunction func(psr *psruntime.PSRuntime) error

func openAll(psr *psruntime.PSRuntime, openFunctions []OpenFunction) error {
	for _, open := range openFunctions {
		if err := open(psr); err != nil {
			return err
		}
	}
	return nil
}
