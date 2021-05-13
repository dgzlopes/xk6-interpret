package interpret

import (
	"go.k6.io/k6/js/modules"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func init() {
	modules.Register("k6/x/interpret", new(Interpret))
}

// Interpret is the k6 interpreted extension.
type Interpret struct {
}

// Run runs a piece of go code.
func (r *Interpret) Run(src string, args interface{}) interface{} {
	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("interpret.Run")
	if err != nil {
		panic(err)
	}

	run := v.Interface().(func(interface{}) interface{})

	return run(args)
}
