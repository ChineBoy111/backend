package middleware

import (
	"bronya.com/gin-gorm/src/api"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/util"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

// Authorize token 鉴权中间件
func Authorize() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		//! Authorization: Bearer <token>
		if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
			api.Err(ctx, api.Resp{Msg: "Token error"})
			return
		}
		token := bearerToken[7:]
		payload, err := util.ParseToken(token)
		if err != nil {
			api.Err(ctx, api.Resp{Msg: err.Error()})
			return
		}
		tokenId := "token" + strconv.Itoa(int(payload.Id))
		cachedToken, err := global.RedisCli.Get(context.Background(), tokenId).Result()
		if err != nil {
			api.Err(ctx, api.Resp{Msg: err.Error()})
			return
		}
		if token != cachedToken {
			api.Err(ctx, api.Resp{Msg: "Token error"})
			return
		}
		//! token 是否过期
		remainTTL, err := global.RedisCli.TTL(context.Background(), token).Result()
		if err != nil {
			api.Err(ctx, api.Resp{Msg: err.Error()})
			return
		}
		if remainTTL < 0 {
			api.Err(ctx, api.Resp{Msg: "token expired"})
			return
		}
		expire := viper.GetDuration("redis.expire") * time.Second
		//! token 续签
		if remainTTL > expire/2 {
			ctx.Next() // 放行
		}
		//! 更新 redis 缓存的 token
		var newToken string
		newToken, err = util.GenToken(payload.Id, payload.Username)
		if err == nil {
			err = global.RedisCli.Set(context.Background(), tokenId, newToken, expire).Err()
		}
		if err != nil {
			api.Err(ctx, api.Resp{Msg: err.Error()})
			return
		}
		ctx.Header("Authorization", "Bearer "+newToken)
		//! 更新 redis 缓存的 loginInfo
		userId := "user" + strconv.Itoa(int(payload.Id))
		loginInfo, err := json.Marshal(map[string]any{
			"id":       payload.Id,
			"username": payload.Username,
		})
		global.RedisCli.Set(context.Background(), userId, loginInfo, expire)
		ctx.Next() // 放行
	}
}
