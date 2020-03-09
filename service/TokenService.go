package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"main/controller"
	"strconv"
	"strings"
	"time"
)

const secret = "rip-key"
const exp = 31536000

func EncodeToken(username string) string {
	// 过期时间
	expireTime := time.Now().Unix() + exp
	// 拼接 payload
	data := username + "." + strconv.FormatInt(expireTime, 10)
	// base64 编码
	base64Data := base64.StdEncoding.EncodeToString([]byte(data))
	// 生成签名
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	// 生成 token
	token := expectedMAC + base64Data
	return token
}

func VerifyToken(c *gin.Context) bool {
	// 提取签名与编码后的 payload
	token := c.GetHeader("Authorization")
	// 验证 Token 长度
	if len(token) < 64 {
		return false
	}
	messageMAC := token[:64]
	base64data := token[64:]
	// base64 解码
	data, err := base64.StdEncoding.DecodeString(base64data)
	if err != nil {
		return false
	}
	// 验证签名
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(data)
	if messageMAC == hex.EncodeToString(mac.Sum(nil)) {
		// 验证用户名与超时时间
		username := strings.Split(string(data), ".")[0]
		expireTime, _ := strconv.ParseInt(strings.Split(string(data), ".")[1], 10, 64)
		if controller.HavingUser(username) && expireTime > time.Now().Unix() {
			return true
		}
		return false
	}
	return false
}

func GetTokenUser(c *gin.Context, username *string) bool {
	// 提取签名与编码后的 payload
	token := c.GetHeader("Authorization")
	// 验证 Token 长度
	if len(token) < 64 {
		return false
	}
	messageMAC := token[:64]
	base64data := token[64:]
	// base64 解码
	data, err := base64.StdEncoding.DecodeString(base64data)
	if err != nil {
		return false
	}
	// 验证签名
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(data)
	if messageMAC == hex.EncodeToString(mac.Sum(nil)) {
		// 验证用户名与超时时间
		*username = strings.Split(string(data), ".")[0]
		expireTime, _ := strconv.ParseInt(strings.Split(string(data), ".")[1], 10, 64)
		if controller.HavingUser(*username) && expireTime > time.Now().Unix() {
			return true
		}
		return false
	}
	return false
}