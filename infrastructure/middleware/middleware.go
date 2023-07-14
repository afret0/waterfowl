package middleware

import "strings"

func MTem(svr string) string {
	t := `package middleware

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"time"
	"sample/source/log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Request.Header.Get("_uid")
		if uid != "" {
			return
		}
		tokenReq := c.Request.Header.Get("token")
		if len(tokenReq) < 1 {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "未携带 token, 请先登录"})
			c.Abort()
			return
		}
		claims, err := GetJWT().ParseToken(tokenReq)
		if err != nil {
			log.GetLogger().Errorln(tokenReq, " ", err)
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
			c.Abort()
			return
		}
		tokenStored := GetTokenManager().GetToken(c, claims.Uid)
		if tokenStored != tokenReq {
			err = errors.New("token 不一致")
			log.GetLogger().Infoln(tokenReq, " ", err.Error())
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
			c.Abort()
			return
		}
		c.Request.Header.Set("_uid", claims.Uid)
		c.Set("claims", claims)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startT := time.Now()
		req, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(req))
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		endT := time.Now()
		latencyT := endT.Sub(startT)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		clientIP := c.ClientIP()
		uid := c.Request.Header.Get("_uid")
		log.GetMiddleWareLogger().WithFields(logrus.Fields{
			"reqTime":  startT.Format("2006-01-02 15:04:05"),
			"latencyT": latencyT.Milliseconds(),
			"method":   reqMethod,
			"uri":      reqUri,
			"clientIP": clientIP,
			"req":      string(req),
			"res":      blw.body.String(),
			"uid":      uid,
		}).Info("请求日志")
	}
}


`
	t = strings.ReplaceAll(t, "sample", svr)
	return t
}
