package auth
import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/wuyan94zl/api/database"
	"github.com/wuyan94zl/api/models"
	"time"
)
// 定义授权保存信息
type CustomClaims struct {
	Id	int
	ExpTime int64
	jwt.StandardClaims
}
// 私钥
const (
	SECRETKEY = "wuyan-secretkey"
)
// 获取用户token值
func GetToken(data *models.User) (map[string]interface{},error) {
	// 7200秒过期
	maxAge := 7200
	expTime := time.Now().Add(time.Duration(maxAge)*time.Second).Unix()
	customClaims := &CustomClaims{
		Id: data.Id,
		ExpTime: expTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime, // 过期时间，必须设置
			Issuer:    "wuyan", // 非必须，也可以填充用户名，
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		return nil,err
	}
	rlt := make(map[string]interface{})
	rlt["expTime"] = expTime
	rlt["token"] = tokenString
	return rlt,nil
}

// 使用token换取user信息
func GetUser(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := int(claims["Id"].(float64))
		user := models.User{}
		database.DB.First(&user,id)
		return user,nil
	} else {
		return nil, err
	}
}