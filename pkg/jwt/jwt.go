package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/wuyan94zl/api/pkg/config"
	"strconv"
	"strings"
	"time"
)

var secretary = config.GetString("jwt.secretary")

type Jwt struct {
	Id        uint64    `json:"id"gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `json:"created_at"gorm:"column:created_at;index"`
	UpdatedAt time.Time `json:"updated_at"gorm:"column:updated_at"`
}

// 定义授权保存信息
type CustomClaims struct {
	Id      uint64
	ExpTime int64
	jwt.StandardClaims
}

// 生成Token
func (user *Jwt) Token() (map[string]interface{}, error) {
	// 7200秒过期
	maxAge, _ := strconv.Atoi(config.GetString("jwt.export"))
	expTime := time.Now().Add(time.Duration(maxAge) * time.Second).Unix()
	customClaims := &CustomClaims{
		Id:      user.Id,
		ExpTime: expTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime, // 过期时间，必须设置
			Issuer:    "wuyan", // 非必须，也可以填充用户名，
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(secretary))
	if err != nil {
		return nil, err
	}
	rlt := make(map[string]interface{})
	rlt["expTime"] = expTime
	rlt["token"] = tokenString
	return rlt, nil
}

// 认证Token
func (user *Jwt) AuthToken(tokenString string) (uint64, interface{}) {
	if tokenString == "" {
		return 0, "认证失败"
	}
	kv := strings.Split(tokenString, " ")
	if kv[0] != "Bearer" {
		return 0, "认证失败"
	}
	tokenString = kv[1]
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretary), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := uint64(claims["Id"].(float64))
		return id, nil
	} else {
		return 0, "认证已过期"
	}
}
