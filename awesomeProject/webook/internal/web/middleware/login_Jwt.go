package middleware

import (
	"awesomeProject/webook/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

type LoginMinddlewareBuilderJwt struct {
}

func (m *LoginMinddlewareBuilderJwt) CheckLoginJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/lonigjwt" {
			return
		}
		//根据约定token在Authorization 头部
		authCode := ctx.GetHeader("Authorization")
		if authCode == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			println("多了空格")
			return
		}
		segs := strings.Split(authCode, " ")
		if len(segs) != 2 {
			//说明Authorization内容乱传的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := segs[1]

		var uc web.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return web.JWTkey, nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if uc.UserAgent != ctx.GetHeader("User-Agent") {
			//后期这个地方需要埋点
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		expireTime := uc.ExpiresAt
		//剩余过期时间小于50秒刷新
		if expireTime.Sub(time.Now()) < time.Second*50 {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenStr, err = token.SignedString(web.JWTkey)
			ctx.Header("x-jwt-token", tokenStr)
			if err != nil {
				log.Println(err)
			}
		}
		ctx.Set("user", uc)

	}
}
