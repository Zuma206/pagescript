package psruntime

import (
	"io"
	"os"

	"github.com/Zuma206/pagescript/eventloop"
	"github.com/Zuma206/pagescript/options"
	"github.com/dop251/goja"
)

func NewPSRuntime(opts ...options.Option[*PSRuntime]) *PSRuntime {
	scanHandlers := newScanHandlers()
	evalHandlers := newEvalHandlers()
	runtime := &PSRuntime{
		passes:       []NodeHandlers{scanHandlers, evalHandlers},
		eventloop:    eventloop.NewEventloop(),
		evalHandlers: evalHandlers,
		scanHandlers: scanHandlers,
		engine:       goja.New(),
		log:          os.Stdout,
	}
	if err := options.Apply(runtime, opts); err != nil {
		return nil
	}
	return runtime
}

var WithLog = options.New(func(psr *PSRuntime, log io.Writer) error {
	psr.log = log
	return nil
})
