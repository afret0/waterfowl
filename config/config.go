package config

func ConfigDevTem() string {
	t := `
{
  "config": "dev",
  "service": {
    "port": 1000
  },
  "mongo": "mongodb://root:****@127.0.0.1:3717/admin?replicaSet=mgset-67171023",
  "userRedis": {
    "addr": "127.0.0.1:6379",
    "user": "default",
    "password": "****",
    "DB": 0
  },
  "redis": {
    "addr": "127.0.0.1:6379",
    "user": "default",
    "password": "****",
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
  "mongo": "mongodb://root:****@127.0.0.1:3717?replicaSet=mgset-67171023&readPreference=secondaryPreferred&MaxPoolSize=200&maxStalenessSeconds=120&heartbeatIntervalMs=20000&maxIdleTimeMs=60000",
  "userRedis": {
    "addr": "127.0.0.1:6379",
    "user": "default",
    "password": "****",
    "DB": 0
  },
  "redis": {
    "addr": "127.0.0.1:6379",
    "user": "default",
    "password": "****",
    "DB": 0
  },
  "rabbitmq": {
    "addr": "amqp://admin:admin@127.0.0.1:5672"
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
  "mongo": "mongodb://root:****@127.0.0.1:3717",
  "rabbitmq": {
    "addr": "amqp://127.0.0.1:5672"
  },
  "userRedis": {
    "addr": "127.0.0.1:6379",
    "user": "default",
    "password": "****",
    "DB": 0
  },
  "redis": {
    "addr": "****",
    "user": "default",
    "password": "qpzm2745",
    "DB": 0
  }
}


`
}
