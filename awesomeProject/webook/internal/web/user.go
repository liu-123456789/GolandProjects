package web

import (
	"awesomeProject/webook/internal/domain"
	"awesomeProject/webook/internal/service"
	"github.com/gin-contrib/sessions"
	"time"

	//jwt
	"github.com/golang-jwt/jwt/v5"
	//正则校验包
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"

	"net/http"
)

const (
	emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	passwordRegexExp  = `^[a-zA-Z0-9_\.\&\@]{6,16}$`
)

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            *service.UserServer
}

func NewUserHandler(svc *service.UserServer) *UserHandler {
	return &UserHandler{
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None), //邮箱格式校验
		passwordRexExp: regexp.MustCompile(passwordRegexExp, regexp.None),
		svc:            svc,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	//没有分组
	//server.POST("/users/signup", h.SignUp)
	//server.POST("/users/login", h.Login)
	//server.POST("/users/edit", h.Edit)
	//server.POST("/users/profile", h.Profile)
	//gin分组路由 Group
	ug := server.Group("/users")
	ug.POST("/signup", h.SignUp)
	ug.POST("/login", h.Login)
	ug.POST("/edit", h.Edit)
	ug.GET("/profile", h.Profile)
	ug.POST("/lonigjwt", h.LoginJwt)
}

// 注册
func (h *UserHandler) SignUp(ctx *gin.Context) {
	type SignupRq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"ConfirmPassword"`
	}

	var req SignupRq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	//isEmail, err := regexp.Match(emailRegexPattern, []byte(req.Email))
	//if err != nil {
	//	ctx.String(http.StatusOK, "系统错误")
	//	return
	//}
	//邮箱格式校验
	isEmail, err := h.emailRexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱错误")
		return
	}

	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次密码不一致")
		return
	}

	ispassword, errs := h.passwordRexExp.MatchString(req.Password)
	if errs != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	//"管理员密码由6-16位组成，可以是数字、字母和.、&、@"
	if !ispassword {
		ctx.String(http.StatusOK, "密码格式错误")
		return
	}
	err = h.svc.Signup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	//if err != nil {
	//	ctx.String(http.StatusOK, "系统错误")
	//	return
	//}
	switch err {
	case nil:
		ctx.String(http.StatusOK, "注册成功")
	case service.ErrDuplicateErr:
		ctx.String(http.StatusOK, "邮箱冲突")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
	//ctx.String(http.StatusOK, "注册成功")
}

// 登录
func (h *UserHandler) Login(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}

	u, err := h.svc.Login(ctx, req.Email, req.Password)
	//if err != nil {
	//	ctx.String(http.StatusOK, "系统错误")
	//	return
	//}
	//ctx.String(http.StatusOK, "登录成功")
	switch err {
	case nil:
		sess := sessions.Default(ctx)
		sess.Set("userId", u.Id)
		sess.Options(sessions.Options{
			//十五分钟
			MaxAge: 900,
			//HttpOnly: true,
		})
		err = sess.Save()
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
		}
		ctx.String(http.StatusOK, "登录成功")
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或者密码不对")
	default:
		ctx.String(http.StatusOK, "系统错误")

	}
}

// 编辑
func (h *UserHandler) Edit(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	println(sess, "100000")
	type RepEdit struct {
		Nickname string `json:"Nickname"`
		Birthday int    `json:"Birthday"`
		About_me string `json:"About_me"`
	}
	var req RepEdit
	err := ctx.Bind(&req)
	if err != nil {
		return
	}

	Uid := sess.Get("userId").(int64)

	if Uid == 0 {
		ctx.String(http.StatusOK, "没有登录")
		return
	}

	err = h.svc.Edit(ctx, domain.User{
		Id:       Uid,
		Nickname: req.Nickname,
		Birthday: req.Birthday,
		About_me: req.About_me,
	})
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.String(http.StatusOK, "修改成功")
}

// 查询
func (h *UserHandler) Profile(ctx *gin.Context) {
	//sess := sessions.Default(ctx)
	//Uid := sess.Get("userId").(int64)
	//通过jwt查看的userkey查看userClaims的uid
	userClaims := ctx.MustGet("user").(UserClaims)
	Uid := userClaims.Uid

	u, err := h.svc.Profile(ctx, Uid)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	type User struct {
		Email    string
		Ctime    int64
		Nickname string
		Birthday int
		About_me string
	}

	ctx.JSON(http.StatusOK, User{
		Email:    u.Email,
		Ctime:    u.Ctime.Unix(),
		Nickname: u.Nickname,
		Birthday: u.Birthday,
		About_me: u.About_me,
	})

}

func (h *UserHandler) LoginJwt(ctx *gin.Context) {
	type ReqLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req ReqLogin
	err := ctx.Bind(&req)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
	}
	//_, err = h.svc.LoginJwt(ctx, req.Email, req.Password)
	u, err := h.svc.LoginJwt(ctx, req.Email, req.Password)
	switch err {
	case nil:
		uc := UserClaims{
			Uid: u.Id,

			//获取userAgent
			UserAgent: ctx.GetHeader("User-Agent"),
			RegisteredClaims: jwt.RegisteredClaims{
				//token的过期时间
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
			},
		}
		// JWT 签名的算法
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)

		tokenStr, err := token.SignedString([]byte(JWTkey))

		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
		}
		ctx.Header("x-jwt-token", tokenStr)
		ctx.String(http.StatusOK, "登录成功")
	case service.AccountDoesNotExist:
		ctx.String(http.StatusOK, "邮箱不存在")
	case service.WrongpPssword:
		ctx.String(http.StatusOK, "密码不对")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}

}

var JWTkey = []byte("qweqweq123123qweqwe")

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       int64
	UserAgent string
}
