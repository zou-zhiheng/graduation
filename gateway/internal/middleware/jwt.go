package middleware

import (
	"errors"
	"fmt"
	"gateway/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type JWtToken struct {
}

type CustomClaims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

func NewJwtToken() *JWtToken {
	return &JWtToken{}
}

var SigningKey = []byte("zouzhiheng-github")

func (j *JWtToken) CreateToken(user model.User) (string, error) {
	//获取token，前两部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{User: user,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(), //签名生效时间
			//ExpiresAt: time.Now().Unix() + 60*60*2, //2小时过期
			Issuer: "zzh", //签发人，
		},
	})
	fmt.Println(token)
	//根据密钥生成加密token，token完整三部分
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return tokenString, err

}
func (j *JWtToken) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("token无效")
	}
	return nil, errors.New("token无效")
}
func (j *JWtToken) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		token := c.Request.Header.Get("token")
		if token == "" {
			//终止
			c.Abort()
			return
		}
		claims, err := j.ParseToken(token)
		if err != nil {
			//终止
			c.Abort()
			return
		}
		//将用户信息储存再上下文
		c.Set("user", claims.User)
		//继续下面的操作
		c.Next()
	}
}
