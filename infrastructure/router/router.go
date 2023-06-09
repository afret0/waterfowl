package router

import "strings"

func InfraRouterTem(svr string) string {
	t := `
package router

import (
	"github.com/gin-gonic/gin"
	sErr "sample/infrastructure/err"
)

var router *Router

type HandleFuncWrap func(c *gin.Context) (interface{}, error)

type Router struct {
	group      map[string]*Group
	e          *gin.Engine
	rootHandle map[string]HandleFuncWrap
}

type item struct {
	method string
	handle HandleFuncWrap
	path   string
}

type serviceResp struct {
	Code int         ` + "`" + "json:" + `"` + "code" + `"` + "`" + `
	Msg  string      ` + "`" + "json:" + `"` + "message" + `"` + "`" + `
	Data interface{} ` + "`" + "json:" + `"` + "data" + `"` + "`" + `
}

func GetRouter(e *gin.Engine) *Router {
	if router != nil {
		return router
	}

	router = new(Router)
	router.group = make(map[string]*Group, 0)
	router.e = e
	router.rootHandle = make(map[string]HandleFuncWrap, 0)

	return router
}

func (r *Router) Group(relativePath string) *Group {
	if _, ok := r.group[relativePath]; ok {
		return r.group[relativePath]
	}

	g := new(Group)
	g.path = relativePath
	g.router = make(map[string]*item, 0)

	mu.Lock()
	defer mu.Unlock()
	r.group[relativePath] = g

	return g
}

func (r *Router) rootGroup() *Group {
	return r.Group("/")
}

func (r *Router) POST(relativePath string, handle HandleFuncWrap) {
	r.rootGroup().POST(relativePath, handle)
}

func (r *Router) GET(relativePath string, handle HandleFuncWrap) {
	r.rootGroup().GET(relativePath, handle)
}

func (r *Router) Use(middleware ...gin.HandlerFunc) {
	r.rootGroup().Use(middleware...)
}

func (r *Router) RegisterRouter(e *gin.Engine) {
	for _, g := range r.group {
		for _, i := range g.router {
			if g.path == "/" {
				r.rootHandle[i.path] = i.handle
			} else {
				r.rootHandle[g.path+i.path] = i.handle
			}
		}
	}
	r.registerRouter(e)
}

func (r *Router) registerRouter(e *gin.Engine) {
	for _, g := range r.group {
		group := e.Group(g.path)
		group.Use(g.use...)
		for _, i := range g.router {
			group.Handle(i.method, i.path, func(ctx *gin.Context) {
				handle := r.rootHandle[ctx.Request.URL.Path]
				resp, err := handle(ctx)
				sr := new(serviceResp)
				sr.Data = resp
				sr.Code = 1
				sr.Msg = "succeed"
				if err != nil {
					sr.Code = 0
					sr.Msg = err.Error()
					errs := sErr.GetErrs(err)
					if errs != nil {
						sr.Code = errs.Code
						sr.Msg = errs.Message
					}
				}
				if sr.Data == nil {
					sr.Data = make(map[string]interface{})
				}
				ctx.JSON(200, sr)
			})
		}
	}
}


`
	t = strings.ReplaceAll(t, "sample", svr)
	return t
}
