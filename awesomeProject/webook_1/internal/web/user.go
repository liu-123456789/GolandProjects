package web

import (
	"awesomeProject/webook_1/internal/domain"
	"awesomeProject/webook_1/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            *service.UserService
}

func NewUserHandler(scv *service.UserService) *UserHandler {
	return &UserHandler{
		emailRexExp:    regexp.MustCompile(emailRegexPatter, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPatter, regexp.None),
		svc:            scv,
	}

}

const (
	emailRegexPatter    = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	passwordRegexPatter = `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`
)

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ur := server.Group("/user")
	ur.POST("/signup", h.signUp)
	//ur.POST("/Login", h.Login)
	ur.POST("/LoginJwt", h.LoginJWT)
	ur.POST("/Edit", h.Edit)
}

func (h *UserHandler) signUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq

	if err := ctx.Bind(&req); err != nil {
		return
	}
	isEmail, err := h.emailRexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return

	}

	if !isEmail {
		ctx.String(http.StatusOK, "邮箱格式错误")
		return
	}

	isPassword, err := h.passwordRexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")

		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密码格式错误")
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次密码不一致")
		return
	}

	err = h.svc.Sigup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	switch err {
	case nil:
		ctx.String(http.StatusOK, "注册成功")
	case service.ErrDuplicateEmail:
		ctx.String(http.StatusOK, "邮箱冲突，请换一个")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}

	//if err != nil {
	//	ctx.String(http.StatusOK, "系统错误")
	//	return
	//}

}

//func (h *UserHandler) Login(ctx *gin.Context) {
//	type LoginReq struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
//	var Req LoginReq
//	err := ctx.Bind(&Req)
//	if err != nil {
//		return
//	}
//
//	u, err := h.svc.Login(ctx, Req.Email, Req.Password)
//	switch err {
//	case nil:
//		sess := sessions.Default(ctx)
//		sess.Set("userId", u.Id)
//		sess.Options(sessions.Options{
//			MaxAge: 10,
//			//HttpOnly: true,
//		})
//		print(sess)
//		err = sess.Save()
//		if err != nil {
//			print(err, "+============")
//			ctx.String(http.StatusOK, "系统错误")
//			return
//		}
//		ctx.String(http.StatusOK, "登录成功")
//	case service.ErrInvalidUserOrPassword:
//		ctx.String(http.StatusOK, "用户名或密码不对")
//	default:
//		ctx.String(http.StatusOK, "系统错误")
//		print(err, "================")
//	}
//
//	//if err != nil {
//	//	ctx.String(http.StatusOK, "系统错误")
//	//	return
//	//}
//
//}

func (h *UserHandler) Edit(ctx *gin.Context) {
	type Req struct {
		Nickname string `json:"Nickname"`
		Birthday string `json:"Birthday"`
		AboutMe  string `json:"AboutMe"`
	}
	var req Req
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	println(req.Birthday, "birthday")
	//birthday, err := time.Parse(time.DateOnly, req.Birthday)
	//if err != nil {
	//	println(err)
	//	ctx.String(http.StatusOK, "生日错误")
	//	return
	//}

	//这里修改使用sessions获取
	//sess := sessions.Default(ctx)
	//Uid := sess.Get("userId").(int64)
	//println(Uid, "=========")
	//if Uid == 0 {
	//	return
	//}

	val, ok := ctx.Get("userId")
	if ok == false {
		return
	}

	Uid := val.(int64)
	println(Uid, "00000")
	if Uid == 0 {
		return
	}

	err = h.svc.UpdateNonSensitiveInfo(ctx, domain.User{
		Id:       Uid,
		Nickname: req.Nickname,
		Birthday: req.Birthday,
		AboutMe:  req.AboutMe,
	})
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}
	ctx.String(http.StatusOK, "修改成功")

}

//	func (h *UserHandler) LoginJwt(ctx *gin.Context) {
//		type LoginJwtReq struct {
//			Email    string `json:"email"`
//			Password string `json:"password"`
//		}
//		var Jwtreq LoginJwtReq
//		err := ctx.Bind(&Jwtreq)
//		if err != nil {
//			return
//		}
//		u, err := h.svc.Login(ctx, Jwtreq.Email, Jwtreq.Password)
//
//		switch err {
//		case nil:
//			uc := UserClaims{
//				Uid: u.Id,
//				RegisteredClaims: jwt.RegisteredClaims{
//
//					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
//				},
//			}
//			token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
//			tokenStr, err := token.SignedString(JWTKEY)
//			if err != nil {
//				ctx.String(http.StatusOK, "登录成功")
//			}
//			ctx.Header("x-jwt-token", tokenStr)
//			ctx.String(http.StatusOK, "登录成功")
//		case service.ErrInvalidUserOrPassword:
//			ctx.String(http.StatusOK, "用户名或密码不对")
//		default:
//			ctx.String(http.StatusOK, "系统错误")
//			print(err, "================")
//
//		}
//
// }
//
// var JWTKEY = []byte("7283o1g852d4dtkz9uj8bu8klalal6u8")
//
// // 声明一个Claims
//
//	type UserClaims struct {
//		jwt.RegisteredClaims
//		Uid int64
//	}
func (h *UserHandler) LoginJWT(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		uc := UserClaims{
			Uid:       u.Id,
			UserAgent: ctx.GetHeader("User-Agent"),
			RegisteredClaims: jwt.RegisteredClaims{
				// 1 分钟过期
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
		tokenStr, err := token.SignedString(JWTKey)
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
		}
		ctx.Header("x-jwt-token", tokenStr)
		ctx.String(http.StatusOK, "登录成功")
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或者密码不对")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

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
	switch err {
	case nil:
		sess := sessions.Default(ctx)
		sess.Set("userId", u.Id)
		sess.Options(sessions.Options{
			// 十分钟
			MaxAge: 30,
		})
		err = sess.Save()
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
			return
		}
		ctx.String(http.StatusOK, "登录成功")
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或者密码不对")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

var JWTKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       int64
	UserAgent string
}
