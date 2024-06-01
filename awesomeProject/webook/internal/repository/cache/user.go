package cache

import (
	"awesomeProject/webook/internal/domain"
	"context"
	"encoding/json"
	"fmt"
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
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	err = json.Unmarshal([]byte(data), &u)
	return u, err
}

func (u UserCache) key(uid int64) string {
	return fmt.Sprintf("user:info:%d", uid)
}

func (c *UserCache) Set(ctx context.Context, du domain.User) error {
	key := c.key(du.Id)
	data, err := json.Marshal(du)
	if err != nil {
		return err
	}
	return c.cmd.Set(ctx, key, data, c.expiration).Err()

}

func NewUserCache(cmd redis.Cmdable) *UserCache {
	return &UserCache{
		cmd: cmd,
		//过期时间15分钟
		expiration: time.Minute * 15,
	}
}
