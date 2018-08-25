package recovery

import (
	"github.com/hlts2/lilty"
)

// Recovery returns recovery middleware for lilty framework
func Recovery() lilty.ChainHandler {
	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			defer func() {
				if err := recover(); err != nil {
					// TODO output log
				}
			}()

			next(ctxt)
		}
	}
}
