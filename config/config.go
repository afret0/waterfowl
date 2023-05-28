package config

func ConfigDevTem() string {
	t := `
{
  "config": "dev",
  "service": {
    "port": 1000
  },
  "mongo": "mongodb://root:WOSHi010@dds-bp1d00511d991bf41229-pub.mongodb.rds.aliyuncs.com:3717,dds-bp1d00511d991bf42154-pub.mongodb.rds.aliyuncs.com:3717/admin?replicaSet=mgset-67171023",
  "userRedis": {
    "addr": "120.27.235.209:6379",
    "user": "default",
    "password": "Qiyiguo2303",
    "DB": 0
  },
  "redis": {
    "addr": "120.27.235.209:6379",
    "user": "default",
    "password": "Qiyiguo2303",
    "DB": 0
  },  
  "rabbitmq": {
    "addr": "amqp://admin:admin@120.27.235.209:5672"
  }
}
`
	return t
}

func ConfigTestTem() string {
	return `
{
  "config": "test",
  "service": {
    "port": 1000
  },
  "mongo": "mongodb://root:WOSHi010@dds-bp1d00511d991bf41.mongodb.rds.aliyuncs.com:3717,dds-bp1d00511d991bf42.mongodb.rds.aliyuncs.com:3717/admin?replicaSet=mgset-67171023&readPreference=secondaryPreferred&MaxPoolSize=200&maxStalenessSeconds=120&heartbeatIntervalMs=20000&maxIdleTimeMs=60000",
  "userRedis": {
    "addr": "120.27.235.209:6379",
    "user": "default",
    "password": "Qiyiguo2303",
    "DB": 0
  },
  "redis": {
    "addr": "120.27.235.209:6379",
    "user": "default",
    "password": "Qiyiguo2303",
    "DB": 0
  },
  "rabbitmq": {
    "addr": "amqp://admin:admin@120.27.235.209:5672"
  }
}
`
}

func ConfigProTem() string {
	return `
{
  "config": "pro",
  "service": {
    "port": 1000
  },
  "mongo": "mongodb://root:Qiyiguo0425@dds-bp19e7cff0b902e41.mongodb.rds.aliyuncs.com:3717,dds-bp19e7cff0b902e42.mongodb.rds.aliyuncs.com:3717/admin?replicaSet=mgset-67861582&readPreference=secondaryPreferred&MaxPoolSize=200&maxStalenessSeconds=120&heartbeatIntervalMs=20000&maxIdleTimeMs=60000",
  "rabbitmq": {
    "addr": "amqp://MjphbXFwLWNuLTlsYjM2eTRmdDAwMjpMVEFJNXQ5VVIyMVJaTXY3dDY4cXVWb2o=:ODlCOEEwQkNFRjIwMDU5QjEzNTMzMUVGMDYyRkIzNDFFQkZBMDREMDoxNjgyNTc1Nzc3MzY0@amqp-cn-9lb36y4ft002.cn-hangzhou.amqp-0.vpc.mq.amqp.aliyuncs.com"
  },
  "userRedis": {
    "addr": "r-bp1630hqxy7sbirzoc.redis.rds.aliyuncs.com:6379",
    "user": "default",
    "password": "Qiyiguo2303",
    "DB": 0
  },
  "redis": {
    "addr": "r-bp1wo8y8i8xdqdxoz7.redis.rds.aliyuncs.com:6379",
    "user": "default",
    "password": "qpzm2745",
    "DB": 0
  }
}


`
}
