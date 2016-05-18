// +build go1.7

package httptreemux

import "net/http"

type ContextMux struct {
	BaseTreeMux
	ContextGroup

	PanicHandler            http.Handler
	NotFoundHandler         http.Handler
	OptionsHandler          http.Handler
	MethodNotAllowedHandler http.Handler
}

func NewContextMux() {
	tm := &ContextMux{
		BaseTreeMux: {
			HeadCanUseGet:          true,
			RedirectTrailingSlash:  true,
			RedirectCleanPath:      true,
			RedirectBehavior:       Redirect301,
			RedirectMethodBehavior: make(map[string]RedirectBehavior),
			PathSource:             RequestURI,
			root:                   &node{path: "/"},
			NotFoundHandler:        http.NotFound,
		},
	}

	tm.BaseTreeMux.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		if tm.PanicHandler != nil {
			r.Context.Params.Set("error", err)
			tm.PanicHandler(w, r)
		}
	}

	tm.BaseTreeMux.OptionsHandler = func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		r.Context.Params.Set("params", params)
		tm.OptionsHandler(w, r)
	}

	tm.BaseTreeMux.MethodNotAllowedHandler = func(w http.ResponseWriter, r *http.Request, methods map[string]HandlerFunc) {
		r.Context.Params.Set("methods", methods)
		tm.MethodNotAllowedHandler(w, r)
	}

	tm.ContextGroup.mux = tm
	return tm
}

type ContextGroup struct {
	Group
}

func (g *ContextGroup) NewGroup(path string) *ContextGroup {
	if len(path) < 1 {
		panic("Group path must not be empty")
	}

	checkPath(path)
	path = g.path + path
	//Don't want trailing slash as all sub-paths start with slash
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return &ContextGroup{path, g.Group.mux}
}

func (g *ContextGroup) Handle(method string, path string, handler http.Handler) {
	g.Group.setHandler(method, path, InternalHandlerFunc(handler))
}

// Syntactic sugar for Handle("GET", path, handler)
func (g *ContextGroup) GET(path string, handler http.Handler) {
	g.Handle("GET", path, handler)
}

// Syntactic sugar for Handle("POST", path, handler)
func (g *ContextGroup) POST(path string, handler http.Handler) {
	g.Handle("POST", path, handler)
}

// Syntactic sugar for Handle("PUT", path, handler)
func (g *ContextGroup) PUT(path string, handler http.Handler) {
	g.Handle("PUT", path, handler)
}

// Syntactic sugar for Handle("DELETE", path, handler)
func (g *ContextGroup) DELETE(path string, handler http.Handler) {
	g.Handle("DELETE", path, handler)
}

// Syntactic sugar for Handle("PATCH", path, handler)
func (g *ContextGroup) PATCH(path string, handler http.Handler) {
	g.Handle("PATCH", path, handler)
}

// Syntactic sugar for Handle("HEAD", path, handler)
func (g *ContextGroup) HEAD(path string, handler http.Handler) {
	g.Handle("HEAD", path, handler)
}

// Syntactic sugar for Handle("OPTIONS", path, handler)
func (g *ContextGroup) OPTIONS(path string, handler http.Handler) {
	g.Handle("OPTIONS", path, handler)
}
