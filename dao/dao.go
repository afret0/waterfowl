package dao

import "strings"

func DaoTem(svr string) string {
	t := `
package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sample/model"
	"sample/source/cache"
	"sample/source/database"
	"sample/source/log"
	"sample/source/tool"
)

var dao *Dao

type Dao struct {
	collection *mongo.Collection
	logger     *logrus.Logger
	tool       *tool.Tool
	redis      redis.UniversalClient
}

func init() {
	dao = new(Dao)
	dao.collection = database.GetMongoDB().GetDatabase().Collection("sample")
	dao.logger = log.GetLogger()
	dao.tool = tool.GetTool()
	dao.redis = cache.GetRedis().GetClient()
}

func GetDao() *Dao {
	return dao
}


func (d *Dao) Find(ctx context.Context, filter interface{}, opt ...*options.FindOptions) ([]*model.Sample, error) {
	cur, err := d.collection.Find(ctx, filter, opt...)
	if err != nil {
		return nil, err
	}
	samples := make([]*model.Sample, 0)
	for cur.Next(ctx) {
		item := new(model.Sample)
		err = cur.Decode(item)
		if err != nil {
			return nil, err
		}
		samples = append(samples, item)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}

	defer func() {
		_ = cur.Close(ctx)
	}()

	return samples, err
}

func (d *Dao) UpdateOne(ctx context.Context,  filter interface{}, update interface{}, opt ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := d.collection.UpdateOne(ctx, filter, update, opt...)
	return res, err
}

func (d *Dao) FindOne(ctx context.Context, filter interface{}, opt ...*options.FindOneOptions) (*model.Sample, error) {
	one := d.collection.FindOne(ctx, filter, opt...)
	u := new(model.Sample)
	err := one.Decode(u)
	return u, err
}

func (d *Dao) InsertOne(ctx context.Context, doc interface{}, opt ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	one, err := d.collection.InsertOne(ctx, doc, opt...)
	return one, err
}

func (d *Dao)Count(ctx context.Context,filter interface{},opt ...*options.CountOptions) (int64,error) {
	count,err := d.collection.CountDocuments(ctx,filter,opt...)
	return count,err
}

`

	t = strings.ReplaceAll(t, "sample", svr)
	t = strings.ReplaceAll(t, "samples", svr+"s")
	t = strings.ReplaceAll(t, "Sample", strings.Title(svr[:1])+svr[1:])
	return t
}
