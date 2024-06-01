package main

import (
	"awesomeProject/webook_1/internal/dao"
	"awesomeProject/webook_1/internal/repository"
	"awesomeProject/webook_1/internal/service"
	"awesomeProject/webook_1/internal/web"
	"awesomeProject/webook_1/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	//if err != nil {
	//	panic(err)
	//}
	//err = dao.InitTables(db)
	//if err != nil {
	//	panic(err)
	//}
	db := initDB()

	//server := gin.Default()
	//server.Use(func(ctx *gin.Context) {
	//	println("这是一个")
	//}, func(context *gin.Context) {
	//	println("第二个mide")
	//},
	//)
	//
	////c := &web.UserHandler{}
	////c := web.NewUserHandler{}
	////跨域
	//server.Use(cors.New(cors.Config{
	//	AllowCredentials: true,
	//	AllowHeaders:     []string{"Content-Type"},
	//	AllowOriginFunc: func(origin string) bool {
	//		if strings.HasPrefix(origin, "http//:localhost") {
	//			return true
	//		}
	//		return strings.Contains(origin, "yuping_050.com")
	//	},
	//	MaxAge: 12 * time.Hour,
	//}), func(ctx *gin.Context) {
	//	println("这是我得midew")
	//})
	server := intWebServer()

	//login := &middleware.LoginMiddlewareBuilder{}
	//store := cookie.NewStore([]byte("secnet"))
	//
	//server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
	initUserHdl(db, server)

	server.Run(":8080")

}

func initUserHdl(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	hel := web.NewUserHandler(us)
	hel.RegisterRoutes(server)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("123456:root@tcp(localhost:3306)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func intWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(func(ctx *gin.Context) {
		println("这是一个")
	}, func(context *gin.Context) {
		println("第二个mide")
	},
	)

	//c := &web.UserHandler{}
	//c := web.NewUserHandler{}
	//跨域
	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http//:localhost") {
				return true
			}
			return strings.Contains(origin, "yuping_050.com")
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		println("这是我得midew")
	})
	//login := &middleware.LoginMiddlewareBuilder{}
	////存储数据得，也就是你 userid 存得哪里
	////直接存cookie
	////store := cookie.NewStore([]byte("secnet"))
	////基于内存实现
	////store := memstore.NewStore([]byte("7283o1g852d4dtkz9uj8bu8klalal6u8"),
	////	[]byte("2t5urfy2b2qahkoyz0rclrndo6u1k3s0"))
	//store, err := redis.NewStore(16, "tcp", "localhost:6379", "",
	//	[]byte("7283o1g852d4dtkz9uj8bu8klalal6u8"),
	//	[]byte("2t5urfy2b2qahkoyz0rclrndo6u1k3s0"))
	//if err != nil {
	//	panic(err)
	//}
	//server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
	userJWT(server)
	//userSerrion(server)
	return server
}
func userSerrion(server *gin.Engine) {
	login := &middleware.LoginMiddlewareBuilder{}
	//存储数据得，也就是你 userid 存得哪里
	//直接存cookie
	//store := cookie.NewStore([]byte("secnet"))
	//基于内存实现
	//store := memstore.NewStore([]byte("7283o1g852d4dtkz9uj8bu8klalal6u8"),
	//	[]byte("2t5urfy2b2qahkoyz0rclrndo6u1k3s0"))
	store, err := redis.NewStore(16, "tcp", "localhost:6379", "",
		[]byte("7283o1g852d4dtkz9uj8bu8klalal6u8"),
		[]byte("2t5urfy2b2qahkoyz0rclrndo6u1k3s0"))
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())

}

func userJWT(server *gin.Engine) {
	//longin := middleware.LoginJWTMiddlewarBuilder{}
	//server.Use(longin.CheckJWTLogin())
	longin := middleware.LoginJWTMiddlewareBuilder{}
	server.Use(longin.CheckLogin())
}
