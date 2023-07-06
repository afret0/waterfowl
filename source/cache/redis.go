package cache

import "strings"

func RedisTem(svr string) string {
	t := `
package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"sample/source/config"
	"sample/source/log"
	"time"
)

type Redis struct {
	logger *logrus.Logger
	ctx    context.Context
	//client redis.UniversalClient
}

var r *Redis
var client *redis.UniversalClient

func GetRedis() *Redis {
	if r != nil {
		return r
	}

	r = new(Redis)
	r.logger = log.GetLogger()
	r.ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	//r.client = r.GetClient()
	return r
}

func (r *Redis) GetClient() redis.UniversalClient {
	if client != nil {
		return *client
	}
	client = r.newClient()
	return *client
}

func (r *Redis) newClient() redis.UniversalClient {
	r.logger.Infof("redis is starting to connect...")

	addr := config.GetConfig().GetString("redis.addr")
	pwd := config.GetConfig().GetString("redis.password")
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{addr},
		Password: pwd,
	})
	err := r.ping(r.ctx, rdb)
	if err != nil {
		r.logger.Fatalf("redis connection failed, err: %s", err.Error())
	} else {
		r.logger.Infof("redis connection succeed...")
	}
	return rdb
}

func (r *Redis) ping(ctx context.Context, rdb redis.UniversalClient) error {
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.GetLogger().Fatalln(err)
	}else {
		log.GetLogger().Infoln("redis ping succeed...")
	}
	return err
}

func (r *Redis) Ping(ctx context.Context) {
	_ = r.ping(ctx, r.client)
}

func (r *Redis) Close() {
	err := r.GetClient().Close()
	if err != nil {
		r.logger.Errorf("redis close err: %s", err.Error())
	} else {
		r.logger.Infof("redis close succeed...")
	}
}

`
	t = strings.ReplaceAll(t, "sample", svr)
	return t
}
