package userCache

import "strings"

func UserRedisTem(svr string) string {
	t := `
package userCache

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
	client redis.UniversalClient
}

var r *Redis

func GetRedis() *Redis {
	if r != nil {
		return r
	}

	r := new(Redis)
	r.logger = log.GetLogger()
	r.ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	r.client = r.GetClient()
	return r
}

func (r *Redis) GetClient() redis.UniversalClient {
	if r.client != nil {
		return r.client
	}
	r.client = r.newClient()
	return r.client
}

func (r *Redis) newClient() redis.UniversalClient {
	r.logger.Infof("redis is start to connect...")

	addr := config.GetConfig().GetString("userRedis.addr")
	pwd := config.GetConfig().GetString("userRedis.password")
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
	}
	return err
}

func (r *Redis) Ping() {
	_ = r.ping(r.ctx, r.client)
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
