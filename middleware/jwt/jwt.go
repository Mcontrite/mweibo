package jwt

import (
	"mweibo/conf"
	"mweibo/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(conf.Serverconfig.JWTSecretKey)

type Claims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, time.Time, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mweibo",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, expireTime, err
}

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

func RefreshToken(token string) (string, time.Time, error) {
	tokenClaims, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			name := claims.Name
			password := claims.Password
			return GenerateToken(name, password)
		}
	}
	return "", time.Now(), nil
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = utils.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = utils.INVALID_PARAMS
			c.Redirect(301, "/admin/login.html")
			return
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = utils.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = utils.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != utils.SUCCESS {
			c.Redirect(301, "/admin/login.html")
			return
		}
		c.Next()
	}
}
