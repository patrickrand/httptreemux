// +build !go1.7

package httptreemux

type InternalHandlerFunc HandlerFunc
type InternalPanicHandler PanicHandler

func createHandlerWrapper(f HandlerFunc) InternalHandlerFunc {
	return InternalHandlerFunc(f)
}

func createPanicWrapper(f PanicHandler) InternalPanicHandler {
	return InternalPanicHandler(f)
}
