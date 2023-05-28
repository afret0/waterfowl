package main

import (
	"bufio"
	"fmt"
	"github.com/afret0/waterfowl/Err"
	"github.com/afret0/waterfowl/build"
	"github.com/afret0/waterfowl/config"
	"github.com/afret0/waterfowl/dao"
	"github.com/afret0/waterfowl/infrastructure/err"
	"github.com/afret0/waterfowl/infrastructure/middleware"
	router2 "github.com/afret0/waterfowl/infrastructure/router"
	"github.com/afret0/waterfowl/model"
	"github.com/afret0/waterfowl/router"
	"github.com/afret0/waterfowl/service"
	"github.com/afret0/waterfowl/source/_http"
	"github.com/afret0/waterfowl/source/cache"
	"github.com/afret0/waterfowl/source/cache/userCache"
	config2 "github.com/afret0/waterfowl/source/config"
	"github.com/afret0/waterfowl/source/database"
	"github.com/afret0/waterfowl/source/limiter"
	sLog "github.com/afret0/waterfowl/source/log"
	"github.com/afret0/waterfowl/source/rabbitmq"
	"github.com/afret0/waterfowl/source/rabbitmq/broker"
	"github.com/afret0/waterfowl/source/rabbitmq/help"
	"github.com/afret0/waterfowl/source/tool"
	"io"
	"log"
	"os"
	"strings"
)

func GetServiceName() string {
	fmt.Println("please input service name...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if text == "" {
		log.Fatal("service is empty")
	}
	fmt.Printf("service name: %s", text)
	return text
}

func Write(content, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.WriteString(file, content)
	if err != nil {
		log.Fatal(err)
	}
}

func main_tem(service string) string {
	t := `
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tylerb/graceful"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"sample/Infrastructure/middleware"
	"sample/router"
	"sample/source/cache"
	"sample/source/config"
	"sample/source/database"
	"sample/source/log"
	"sample/source/tool"
	"time"
)

func main() {
	logger := log.GetLogger()
	cfg := config.GetConfig()

	logger.Infoln("service Initialize resources...")
	logger.Infof("service use config, current cfg: %s", cfg.Get("config"))

	env := tool.GetTool().GetEnv()
	if env == "pro" {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.Infof("service env: %s", env)

	ctx := context.Background()
	database.GetMongoDB().Ping(ctx)
	cache.GetRedis().Ping(ctx)

	engine := gin.New()
	engine.Use(gin.Recovery(), middleware.LoggerMiddleware())
	router.RegisterRouter(engine)

	logger.Infoln("service is running...")
	port := cfg.GetString("service.port")
	logger.Infof("service port: %s", port)

	srv := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: engine,
		},
		NoSignalHandling: true,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalf("service run failed, err: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Printf("HTTP server shutdown failed: %v", err)
	} else {
		logger.Println("HTTP server shutdown gracefully")
	}

	database.GetMongoDB().Disconnect()
	cache.GetRedis().Close()
	return
}

`
	t = strings.ReplaceAll(t, "sample", service)
	return t
}

func sample_sh_tem(svr string) string {
	t := `
git add ./*
git commit -am "buid at ` + "`" + "date" + "`" + `"
git push
ssh root@dev.kekeyuyin.com "cd /root/sample && git pull && supervisorctl restart sample"
`
	t = strings.ReplaceAll(t, "sample", svr)
	return t
}

func Gitignore_tem() string {
	t := `
	# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

*.sample

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
.idea
.vscode
bin

# Dependency directories (remove the comment below to include it)
# vendor/

`
	return t
}

func NewService() {
	//svr := GetServiceName()
	svr := "sample"

	os.MkdirAll(svr, 0755)
	Write(main_tem(svr), svr+"/main.go")

	os.MkdirAll(svr+"/config", 0755)
	Write(config.ConfigDevTem(), svr+"/config/configDev.json")
	Write(config.ConfigTestTem(), svr+"/config/configTest.json")
	Write(config.ConfigProTem(), svr+"/config/config.json")

	os.MkdirAll(svr+"/model", 0755)
	Write(model.ModelTem(svr), svr+"/model/model.go")

	os.MkdirAll(svr+"/router", 0755)
	Write(router.RouterTem(svr), svr+"/router/router.go")

	os.MkdirAll(svr+"/handler", 0755)
	Write(service.SvrTem(svr), svr+"/handler/service.go")

	os.MkdirAll(svr+"/infrastructure", 0755)
	os.MkdirAll(svr+"/infrastructure/middleware", 0755)
	Write(middleware.MTem(svr), svr+"/infrastructure/middleware/middleware.go")
	Write(middleware.TokenTem(svr), svr+"/infrastructure/middleware/token.go")
	os.MkdirAll(svr+"/infrastructure/err", 0755)
	Write(err.InfrErrTem(), svr+"/infrastructure/err/error.go")
	os.MkdirAll(svr+"/infrastructure/router", 0755)
	Write(router2.InfraRouterTem(svr), svr+"/infrastructure/router/router.go")
	Write(router2.GroupTem(svr), svr+"/infrastructure/router/group.go")

	os.MkdirAll(svr+"/source", 0755)
	os.MkdirAll(svr+"/source/cache", 0755)
	Write(cache.RedisTem(svr), svr+"/source/cache/redis.go")
	os.MkdirAll(svr+"/source/cache/userCache", 0755)
	Write(userCache.UserRedisTem(svr), svr+"/source/cache/userCache/userRedis.go")
	os.MkdirAll(svr+"/source/config", 0755)
	Write(config2.ConfigTem(svr), svr+"/source/config/config.go")
	os.MkdirAll(svr+"/source/database", 0755)
	Write(database.DatabaseTem(svr), svr+"/source/database/db.go")
	os.MkdirAll(svr+"/source/log", 0755)
	Write(sLog.LogTem(svr), svr+"/source/log/logger.go")
	os.MkdirAll(svr+"/source/limiter", 0755)
	Write(limiter.LimiterTem(svr), svr+"/source/limiter/limiter.go")
	os.MkdirAll(svr+"/source/_http", 0755)
	Write(_http.HttpTem(), svr+"/source/_http/request.go")
	os.MkdirAll(svr+"/source/tool", 0755)
	Write(tool.ToolTem(), svr+"/source/tool/tool.go")
	os.MkdirAll(svr+"/source/rabbitmq", 0755)
	Write(rabbitmq.ConsumerTem(svr), svr+"/source/rabbitmq/consumer.go")
	Write(rabbitmq.ProducerTem(svr), svr+"/source/rabbitmq/producer.go")
	os.MkdirAll(svr+"/source/rabbitmq/help", 0755)
	Write(help.HelpTem(svr), svr+"/source/rabbitmq/help/help.go")
	os.MkdirAll(svr+"/source/rabbitmq/broker", 0755)
	Write(broker.BrokerTem(), svr+"/source/rabbitmq/broker/broker.go")
	Write(broker.AmqpTem(svr), svr+"/source/rabbitmq/broker/amqp.go")

	os.MkdirAll(svr+"/"+"Err", 0755)
	Write(Err.ErrTem(svr), svr+"/"+"Err"+"/err.go")

	os.MkdirAll(svr+"/handler", 0755)

	os.MkdirAll(svr+"/dao", 0755)
	Write(dao.DaoTem(svr), svr+"/dao/dao.go")

	Write("#"+svr, svr+"/README.md")

	// os.MkdirAll(svr+"/sample.sh", 0755)
	Write(sample_sh_tem(svr), svr+"/test.sh")

	Write(Gitignore_tem(), svr+"/.gitignore")

	Write(build.BuildSh_tem(), svr+"/build.sh")

	Write(build.DockerFileTem(), svr+"/Dockerfile")

	Write(build.MakeFile_tem(svr), svr+"/Makefile")
}

func main() {
	NewService()
}
