package middleware

import (
	"awesomeProject/webook_1/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

//type LoginJWTMiddlewarBuilder struct {
//}

//	func (m *LoginJWTMiddlewarBuilder) CheckJWTLogin() gin.HandlerFunc {
//		return func(ctx *gin.Context) {
//			path := ctx.Request.URL.Path
//			if path == "/user/signup" || path == "/user/LoginJwt" {
//				return
//			}
//			//根据约定，token在Authorization 头部
//			authCode := ctx.GetHeader("Authorization")
//			if authCode == "" {
//				//没登录 没token
//				ctx.AbortWithStatus(http.StatusUnauthorized)
//				println("======错误4")
//				return
//			}
//
//			sage := strings.Split(authCode, " ")
//			//Authorization 不等了2个字符，内容是乱传的
//			if len(sage) != 2 {
//				ctx.AbortWithStatus(http.StatusUnauthorized)
//				println("======错误3")
//				println(sage)
//				return
//			}
//			tokenStr := sage[1]
//			var uc web.UserClaims
//			token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
//				return web.JWTKEY, nil
//			})
//			//token是伪造的
//			if err != nil {
//				ctx.AbortWithStatus(http.StatusUnauthorized)
//				println("======错误2")
//				return
//			}
//			if token == nil || !token.Valid {
//				//token非法或者过期
//				ctx.AbortWithStatus(http.StatusUnauthorized)
//				println("======错误1")
//				return
//
//			}
//			expireTime := uc.ExpiresAt
//			//token过期了
//			//if expireTime.Before(time.Now()) {
//			//	ctx.AbortWithStatus(http.StatusUnauthorized)
//			//	return
//			//}
//
//			//剩余过期时间<50s过期
//			if expireTime.Sub(time.Now()) < time.Second*50 {
//				uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 5))
//				tokenStr, err = token.SignedString(web.JWTKEY)
//				ctx.Header("x-jwt-token", tokenStr)
//				if err != nil {
//					log.Print(err)
//				}
//
//			}
//			ctx.Set("user", uc)
//
//		}
//	}
type LoginJWTMiddlewareBuilder struct {
}

func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/user/signup" || path == "/user/LoginJwt" {
			// 不需要登录校验
			return
		}
		// 根据约定，token 在 Authorization 头部
		// Bearer XXXX
		authCode := ctx.GetHeader("Authorization")
		if authCode == "" {
			// 没登录，没有 token, Authorization 这个头部都没有
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(authCode, " ")
		if len(segs) != 2 {
			// 没登录，Authorization 中的内容是乱传的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		var uc web.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return web.JWTKey, nil
		})
		if err != nil {
			// token 不对，token 是伪造的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid {
			// token 解析出来了，但是 token 可能是非法的，或者过期了的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if uc.UserAgent != ctx.GetHeader("User-Agent") {
			// 后期我们讲到了监控告警的时候，这个地方要埋点
			// 能够进来这个分支的，大概率是攻击者
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		expireTime := uc.ExpiresAt
		// 不判定都可以
		//if expireTime.Before(time.Now()) {
		//	ctx.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		// 剩余过期时间 < 50s 就要刷新
		if expireTime.Sub(time.Now()) < time.Second*50 {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 5))
			tokenStr, err = token.SignedString(web.JWTKey)
			ctx.Header("x-jwt-token", tokenStr)
			if err != nil {
				// 这边不要中断，因为仅仅是过期时间没有刷新，但是用户是登录了的
				log.Println(err)
			}
		}
		ctx.Set("userId", uc.Uid)
	}
}
