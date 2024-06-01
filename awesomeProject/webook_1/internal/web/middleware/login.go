package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
}

func (m *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	//注册一下这个类型
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/user/signup" || path == "/user/Login" {
			return
		}
		sess := sessions.Default(ctx)
		//sess.Options(sessions.Options{MaxAge: 966})
		userId := sess.Get("userId")
		if userId == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//拿到当前时间搓
		now := time.Now()
		//我怎么知道，要刷新呢？
		//假如我们每分钟刷新一次

		const updateTimeKeg = "update_time"
		val := sess.Get(updateTimeKeg)
		lastUpdateTime, ok := val.(time.Time)
		if val == nil || !ok || now.Sub(lastUpdateTime) > time.Second*10 {
			//第一次进来
			sess.Set(updateTimeKeg, now)
			sess.Set("userId", userId)
			err := sess.Save()
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
