package cache

import (
	"context"
	"fmt"
	"gitee.com/geekbang/basic-go/webook/search/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserCache struct {
	//redis 客户端的引用
	cmd redis.Cmdable
	//过期时间
	expiration time.Duration
}

func (c *UserCache) Get(ctx context.Context, uid int64) (domain.User, error) {
	key := c.key(uid)
	data, err := c.cmd.Get(ctx, key).Result()
}

func (u UserCache) key(uid int64) string {
	return fmt.Sprintf("user.info:%d", uid)
}

func NewUserCache(cmd redis.Cmdable) *UserCache {
	return &UserCache{
		cmd: cmd,
		//过期时间15分钟
		expiration: time.Minute * 15,
	}
}
