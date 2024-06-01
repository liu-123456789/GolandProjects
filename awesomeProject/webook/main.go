package main

import (
	"awesomeProject/webook/internal/repository"
	"awesomeProject/webook/internal/repository/dao"
	"awesomeProject/webook/internal/service"
	"awesomeProject/webook/internal/web"
	"awesomeProject/webook/internal/web/middleware"
	"awesomeProject/webook/pkg/ginx/middleware/ratelimit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

// 集中式注册
func main() {

	//db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/webook"))
	//if err != nil {
	//	panic(err)
	//}

	//db := initDB()
	//建表语句
	//dao.InitTables(db)
	//server := initWebServer()
	//
	//initUserHdl(db, server)
	server := gin.Default()
	server.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello,启动成功")
	})
	server.Run(":8089")

	//hdl 等于指向web包的UserHandler的类型
	//hdl := &web.UserHandler{}
	//hdl 等于指向web包的NewUserHandler的类型

	//ud := dao.NewUserdao(db)
	//ur := repository.NewUserRepository(ud)
	//us := service.NewUserServer(ur)
	//hdl := web.NewUserHandler(us)
	//hdl.RegisterRoutes(server)
	//server.Run(":8089")
}

func userJwt(server *gin.Engine) {
	login := middleware.LoginMinddlewareBuilderJwt{}
	server.Use(login.CheckLoginJwt())
}

func UserSession(server *gin.Engine) {
	login := &middleware.LoginMinddlewareBuilder{}
	//直接存cookie
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
}

// 链接mysql
func initDB() *gorm.DB {
	//gorm 连接mysql
	db, err := gorm.Open(mysql.Open(("root:123456@tcp(localhost:3306)/webook")))
	//出现错误打印出来
	if err != nil {
		panic(err)
	}

	//建表语句
	errs := dao.InitTables(db)
	if errs != nil {
		panic(errs)
	}
	//没有错误就返回出去
	return db
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	//跨域问题
	server.Use(cors.New(cors.Config{
		AllowCredentials: true,                                          //应许Cookie带过来
		AllowHeaders:     []string{"application/json", "Authorization"}, //允许请求得头部带参数过来
		ExposeHeaders:    []string{"x-jwt-token"},                       //允许前端访问带你得后端响应头部过来
		//AllowOriginFunc 那些来源可以允许
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			//允许公司得域名
			//if strings.Contains(origin, "公司域名") {
			//	return true
			//}
			return strings.Contains(origin, "http://127.0.0.1")
		},
		MaxAge: 12 * time.Hour, //检测多长时间
	}), func(context *gin.Context) {
		println("这是我的mdw")

	})

	//redis 连接客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	//限流,一秒钟，1个请求
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 1).Build())

	userJwt(server)
	//UserSession(server)
	return server
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {

	ud := dao.NewUserdao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserServer(ur)
	hdl := web.NewUserHandler(us)
	hdl.RegisterRoutes(server)

}
