package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 自定义JWT声明
type CustomClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

const (
	JWTSecretKey = "peihuachen2001@gmail.com"
)

// 生成随机盐值
func GenerateSalt(length int) string {
	return grand.S(length, false)
}

// 密码加密 (双重MD5)
func EncryptPassword(password, salt string) string {
	// 加密(加密密码 + 加密盐)
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

//// 生成Token
//func GenerateToken(ctx context.Context, userId int) (string, time.Time, error) {
//	// 获取JWT配置
//	jwtSecret := g.Cfg().MustGet(ctx, "jwt.secret").String()
//	expireDays := g.Cfg().MustGet(ctx, "jwt.expire", 7).Int()
//
//	// 计算过期时间
//	expireTime := time.Now().Add(time.Duration(expireDays) * 24 * time.Hour)
//
//	// 创建Token
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"user_id": userId,
//		"exp":     expireTime.Unix(),
//	})
//
//	// 签名Token
//	signedToken, err := token.SignedString([]byte(jwtSecret))
//	if err != nil {
//		g.Log().Errorf(ctx, "生成Token失败: %v", err)
//		return "", time.Time{}, errors.New("系统错误")
//	}
//
//	return signedToken, expireTime, nil
//}

// 生成JWT Token
func GenerateToken(userId int) (string, time.Time, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), //生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", time.Time{}, err
	}
	return signedToken, expireTime, nil
}

// 解析JWT Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	//三个参数分别为：要解析的JWT令牌字符串、目标声明的结构体指针、密钥提供函数
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
