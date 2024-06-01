package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMinddlewareBuilder struct {
}

func (m *LoginMinddlewareBuilder) CheckLogin() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			//不需要校验
			return
		}
		sess := sessions.Default(context)
		userId := sess.Get("userId")
		if userId == nil {
			//中断不用执行后面业务逻辑了
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		now := time.Now()
		const updatetTimeKey = "update_time"
		val := sess.Get(updatetTimeKey)

		if val == nil {
			sess.Set(updatetTimeKey, now)
			sess.Set("userId", userId)
			err := sess.Save()
			if err != nil {
				fmt.Println(err)
			}
		}
		lastUpdateTime, ok := val.(time.Time)
		if !ok {
			sess.Set(updatetTimeKey, now)
			sess.Set("userId", userId)
			err := sess.Save()
			if err != nil {
				fmt.Println(err)
			}
		}
		//过去了一分钟
		if now.Sub(lastUpdateTime) > time.Minute {
			sess.Set(updatetTimeKey, now)
			sess.Set("userId", userId)
			err := sess.Save()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
