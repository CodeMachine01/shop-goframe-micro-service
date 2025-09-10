package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"shop-goframe-micro-service/utility"
)

const (
	CtxUserId gctx.StrKey = "userId"
)

func JWTAuth(r *ghttp.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "未提供Token"))
		return
	}

	//移除Bearer前缀"Bearer "（注意后面有个空格）
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	claims, err := utility.ParseToken(tokenString)
	if err != nil || claims == nil {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "无效Token"))
		return
	}

	//将用户ID注入上下文
	r.SetCtxVar(CtxUserId, claims.UserId)
	//调用下一个中间件或最终处理程序
	r.Middleware.Next()
}
