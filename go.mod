module github.com/kiebitz-oss/services

go 1.16

require (
	github.com/bsm/redislock v0.7.1
	github.com/go-redis/redis/v8 v8.11.4
	github.com/kiprotect/go-helpers v0.0.0-20210706144641-b74c3f0f016d
	github.com/prometheus/client_golang v1.11.1
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli v1.22.5
)

// replace github.com/kiprotect/go-helpers => ../../../geordi/kiprotect/go-helpers
