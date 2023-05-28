package database

import "strings"

func DatabaseTem(svr string) string {
	t := `
package database

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sample/source/config"
	"sample/source/log"
	"time"
)

type MongoDB struct {
	logger *logrus.Logger
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

var m *MongoDB

func GetMongoDB() *MongoDB {
	if m != nil {
		return m
	}

	m = new(MongoDB)
	m.logger = log.GetLogger()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	m.ctx = ctx
	m.client = m.newClient(m.ctx)

	return m
}

func (m *MongoDB) newClient(ctx context.Context) *mongo.Client {
	m.logger.Infoln("MongoDB is starting to connect...")

	uri := config.GetConfig().GetString("mongo")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20))
	if err != nil {
		m.logger.Fatalf("MongoDB connection failed, err: %s, uri: %s", err.Error(), uri)
	} else {
		m.logger.Infof("MongoDB connection succeed...")
	}
	return client
}

func (m *MongoDB) Ping(ctx context.Context) {
	err := m.client.Ping(ctx, readpref.Primary())
	if err != nil {
		m.logger.Fatalf("mongoDB ping err: %s", err.Error())
	} else {
		m.logger.Info("mongoDB ping succeed...")
	}
}

func (m *MongoDB) GetDatabase(name ...string) *mongo.Database {
	if m.db == nil {
		m.db = m.client.Database("guoguo")
		if len(name) > 0 {
			m.db = m.client.Database(name[0])
		}
	}
	return m.db
}

func (m *MongoDB) GetCollection(col string) *mongo.Collection {
	if m.db == nil {
		m.db = m.GetDatabase()
	}
	return m.db.Collection(col)
}

func (m *MongoDB) Disconnect() {
	err := m.client.Disconnect(m.ctx)
	if err != nil {
		m.logger.Fatalf("mongoDB disconnect err: %s", err.Error())
	} else {
		m.logger.Info("mongoDB disconnect succeed...")
	}
}

`
	t = strings.ReplaceAll(t, "sample", svr)
	return t

}
