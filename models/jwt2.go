package main

import (
	"github.com/Unknwon/com"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"strconv"
	"time"
)

// JWT Payload结构
type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var db1 *gorm.DB

func init() {
	//创建一个数据库的连接
	var err error
	db1, err = gorm.Open("sqlite3", "jwt2.db1")
	if err != nil {
		panic("failed to connect database")
	}

	//迁移the schema
	db1.AutoMigrate(&Users{})
}

// 生成JWT
func GenerateToken(username string, password string) (string, error) {
	expireAt := time.Now().Add(30 * time.Minute)
	issuedBy := "nobody"
	secret := "secret"

	// 定义Payload
	claim := Users{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
			Issuer:    issuedBy,
		},
	}

	// 定义签名算法, 签名, 生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(secret))
	return ss, err
}

// 解析JWT
func ParseToken(ss string) (*Users, error) {
	secret := "secret"

	// 解析Payload
	token, err := jwt.ParseWithClaims(ss, &Users{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// 验证Payload
	if err == nil && token != nil {
		if claim, ok := token.Claims.(*Users); ok && token.Valid {
			return claim, nil
		}
	}
	return nil, err
}

var code int

func loginn(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	token, err := GenerateToken(username, password)
	if err != nil {
		code = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"cookie": token,
	})

}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

func main() {
	r := gin.Default()

	r.POST("/login", loginn)
	r.Run()
}
