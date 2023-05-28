package service

import "strings"

func SvrTem(svr string) string {
	t := `
package handler

import (
	"sample/source/cache"
	"sample/source/database"
	"sample/source/log"
	"sample/source/tool"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var svr *Service

type Service struct {
	logger *logrus.Logger
	tool   *tool.Tool
}

func GetService() *Service {
	if svr != nil {
		return svr
	}
	svr = new(Service)
	svr.logger = log.GetLogger()
	svr.tool = tool.GetTool()

	return svr
}

type PingResp struct {
	Ping string ` + "`" + "json:" + `"` + "ping" + `"` + "`" + `
}

func (s *Service) Ping(ctx *gin.Context) (interface{}, error) {
	database.GetMongoDB().Ping(ctx)
	cache.GetRedis().Ping(ctx)
	return &PingResp{Ping: "pong"}, nil
}


`
	t = strings.ReplaceAll(t, "sample", svr)
	return t
}
