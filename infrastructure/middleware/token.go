package middleware

import (
	"strings"
)

func TokenTem(svr string) string {
	t := `
package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sample/source/cache/userCache"
	"sample/source/database"
	"sample/source/log"
	"time"
)

type User struct {
	Token string ` + "`" + "bson:" + "token json:" + `"` + "token" + "`" + `
}

var j *JWT
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	signKey          []byte = []byte("AllYourBase")
)

var tokenManager *Manager

func GetTokenManager() *Manager {
	if tokenManager == nil {

		tokenManager = new(Manager)
		tokenManager.redis = userCache.GetRedis().GetClient()
		tokenManager.collection = database.GetMongoDB().GetCollection("user")
		tokenManager.logger = log.GetLogger()
	}
	return tokenManager
}

type Manager struct {
	redis      redis.UniversalClient
	collection *mongo.Collection
	logger     *logrus.Logger
}

func (t *Manager) fmtTokenKey(key string) string {
	return fmt.Sprintf("token:%s", key)
}

func (t *Manager) UpdateToken(ctx context.Context, uid, token string) error {
	filter := bson.M{"uid": uid}
	upt := bson.M{"$set": bson.M{"token": token, "updateTime": time.Now().UnixMilli()}}
	opt := new(options.UpdateOptions)
	_, err := t.collection.UpdateOne(ctx, filter, upt, opt)
	t.redis.Set(ctx, t.fmtTokenKey(uid), token, 3600*24*100*time.Second)
	return err
}

func (t *Manager) GetToken(ctx context.Context, uid string) string {
	token := t.redis.Get(ctx, t.fmtTokenKey(uid)).Val()
	if token != "" {
		return token
	}
	filter := bson.M{"uid": uid}
	opt := new(options.FindOneOptions)
	opt.SetProjection(bson.M{"token": 1, "_id": 0})
	one := t.collection.FindOne(ctx, filter, opt)
	if one == nil {
		t.logger.Errorln("uid:  " + uid + "token not found")
		return ""
	}
	u := new(User)
	err := one.Decode(u)
	if err != nil {
		t.logger.Errorln(err)
		return ""
	}
	return u.Token
}

type JWT struct {
	logger *logrus.Logger
}

type Claims struct {
	Uid string
	jwt.StandardClaims
}

func (j *JWT) GenerateToken(uid string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3600 * time.Second * 24 * 1000)
	issuer := "guoguo"
	c := Claims{uid, jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: issuer}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenS, err := token.SignedString(signKey)
	if err != nil {
		j.logger.Errorln(err, "gen token userErr")
	}
	return tokenS, err
}

func (j *JWT) ParseToken(token string) (*Claims, error) {
	c := new(Claims)
	tokenClaims, err := jwt.ParseWithClaims(token, c, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		j.logger.Errorln(token, err)
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GetJWT() *JWT {
	if j == nil {
		j = new(JWT)
		j.logger = log.GetLogger()

	}
	return j
}

`
	t = strings.ReplaceAll(t, "sample", svr)
	return t
}
