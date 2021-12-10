package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AccessTokenExpireDuration  = time.Minute * 10
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

var mySecret = []byte("gin-bluebell-example")

// MyClaims 自定义声明结构体并内嵌 jwt.StandardClaims
// jwt包自带的 jwt.StandardClaims 只包含了官方字段
// 我们这里需要额外记录一个 username 字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaiims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 生成 access token 和 refresh token
func GenToken(userID int64) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaiims{
		userID, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(), // 过期时间
			Issuer:    "gin-bluebell",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	// refresh token 不需要存任何自定义数据
	rToken, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), // 过期时间
		Issuer:    "gin-bluebell",                                    // 签发人
	}).SignedString(mySecret)

	// 使用指定的 secret 签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析 JWT
func ParseToken(tokenString string) (*MyClaiims, error) {
	// 解析token
	var mc = new(MyClaiims)
	token, err := jwt.ParseWithClaims(
		tokenString,
		mc,
		func(token *jwt.Token) (i interface{}, err error) {
			return mySecret, nil
		},
	)
	if err != nil {
		return nil, err
	}
	if token.Valid {
		// 校验 token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新 AccessToken
func RefreshToken(aToken, rToken string) (newAtoken, newRToken string, err error) {
	// refresh token 无效直接返回
	_, err = jwt.Parse(rToken, func(aToken *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return
	}

	// 从旧 access token 中解析出 claims 数据
	var claims MyClaiims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(aToken *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	v, _ := err.(*jwt.ValidationError)

	// 当access token 是过期错误 并且 refresh token 没有过期时就创建一个薪的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID)
	}
	return
}
