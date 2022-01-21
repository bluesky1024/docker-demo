package main

import(
	"github.com/go-redis/redis"
)

const(
	cntKey = "blue_cnt"
)

type CntRepo interface {
	Incr(i int64) error
	Get() (int64,error)
}

type cntRepoImpl struct {
	client *redis.Client
}

func NewCntRepo() CntRepo {
	redisClient := redis.NewClient(&redis.Options{
		Addr:"blue_redis:6379",
	})
	return &cntRepoImpl{
		client: redisClient,
	}
}

func (c *cntRepoImpl) Incr(i int64) error {
	_,err := c.client.IncrBy(cntKey,i).Result()
	return err
}

func (c *cntRepoImpl) Get() (int64,error) {
	return c.client.Get(cntKey).Int64()
}
