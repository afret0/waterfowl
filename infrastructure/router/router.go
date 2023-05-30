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
	rGroup map[string]*Group
	e      *gin.Engine
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
	router.rGroup = make(map[string]*Group, 0)
	router.e = e

	return router
}

func (r *Router) Group(relativePath string) *Group {
	if _, ok := r.rGroup[relativePath]; ok {
		return r.rGroup[relativePath]
	}

	g := new(Group)
	g.path = relativePath
	g.router = make(map[string]*item, 0)

	mu.Lock()
	defer mu.Unlock()
	r.rGroup[relativePath] = g

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

func (r *Router) RegisterRouter() {
	r.registerRouter()
}

func (r *Router) registerRouter() {
	for _, g := range r.rGroup {
		group := r.e.RouterGroup.Group(g.path)
		group.Use(g.use...)
		for _, i := range g.router {
			method := i.method
			group.Handle(method, i.path, func(ctx *gin.Context) {
				resp, err := i.handle(ctx)
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
