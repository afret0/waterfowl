package response

func ResTem() string {
	t := `
package response

import (
	"github.com/gin-gonic/gin"
)

type Manager struct {
}

type Response struct {
	Code    int         ` + "`" + "json:" + `"` + "code" + `"` + "`" + `
	Message string      ` + "`" + "json:" + `"` + "message" + `"` + "`" + `
	Data    interface{} ` + "`" + "json:" + `"` + "data" + `"` + "`" + `
}

var manager *Manager

func init() {
	manager = new(Manager)
}

func (r Manager) SucceedResponse(ctx *gin.Context, data interface{}) {
	succeedResponse := new(Response)
	succeedResponse.Code = 1
	succeedResponse.Message = "succeed"
	if data == nil {
		data = make(map[string]string)
	}
	succeedResponse.Data = data
	ctx.JSON(200, succeedResponse)
}

func (r Manager) SucceedResponseWithoutData(ctx *gin.Context) {
	succeedResponse := new(Response)
	succeedResponse.Code = 1
	succeedResponse.Message = "succeed"
	succeedResponse.Data = make(map[string]string)
	ctx.JSON(200, succeedResponse)
}

func (r Manager) FailedResponse(ctx *gin.Context, err error) {
	s := new(Response)
	
	s.Code = 0
	if len(code) > 0 {
		s.Code = code[0]
	}
	s.Message = msg
	s.Data = make(map[string]string)
	ctx.JSON(200, s)
}

func GetManager() *Manager {
	return manager
}

`
	return t
}
