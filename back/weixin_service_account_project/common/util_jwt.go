package common

import (
	"github.com/dgrijalva/jwt-go"
	"szyx_back/configs"
	"time"
)

var jwtSecret = []byte(configs.JWT_SECRET)

type Claims struct {
	EmpCode  string `json:"empCode"`
	CorpCode string `json:"corpCode"`

	jwt.StandardClaims
}

// 产生token的函数
func GenerateToken(empCode, password string, corpCode string) (string, error) {

	//token 过期时间 access_token 过期时间为5分钟
	nowTime := time.Now()
	expireTime := nowTime.Add(5000 * time.Minute)

	//1. aud 标识token的接收者.
	//2. exp 过期时间.通常与Unix UTC时间做对比过期后token无效
	//3. jti 是自定义的id号
	//4. iat 签名发行时间.
	//5. iss 是签名的发行者.
	//6. nbf 这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
	//7. sub 签名面向的用户
	claims := Claims{
		empCode,
		corpCode,
		jwt.StandardClaims{
			Issuer:    "htfkTeam",
			ExpiresAt: expireTime.Unix(),
		},
	}
	//
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 产生token的函数
func GenerateRefreshToken(empCode, password string, corpCode string) (string, error) {

	//token 过期时间 refresh_token 过期时间为30分钟
	nowTime := time.Now()
	expireTime := nowTime.Add(30 * time.Minute)

	//1. aud 标识token的接收者.
	//2. exp 过期时间.通常与Unix UTC时间做对比过期后token无效
	//3. jti 是自定义的id号
	//4. iat 签名发行时间.
	//5. iss 是签名的发行者.
	//6. nbf 这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
	//7. sub 签名面向的用户
	claims := Claims{
		empCode,
		corpCode,
		jwt.StandardClaims{
			Issuer:    "htfkTeam",
			ExpiresAt: expireTime.Unix(),
		},
	}
	//
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 验证token的函数
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
