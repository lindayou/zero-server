package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateSmsCode(witdh int) string {
	//生成width长度的短信验证码

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < witdh; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// 生成token
type UserClaim struct {
	Id       int
	Identity string
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(id int, identity string, name string, jwtKey string) (string, error) {
	//ID Identity Name
	//定义对象
	uc := new(UserClaim)
	uc.Username = name
	uc.Id = id
	uc.Identity = identity
	uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "token SignedString failed,err : ", err
	}

	return tokenString, nil
}
