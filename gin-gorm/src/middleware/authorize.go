package middleware

import (
	"bronya.com/gin-gorm/src/api"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/util"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

// Authorize token 鉴权中间件
func Authorize() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		//! Authorization: Bearer <token>
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			api.Err(ctx, api.Resp{Msg: "Token error"})
			return
		}
		token = token[7:]
		payload, err := util.ParseToken(token)
		if err != nil {
			api.Err(ctx, api.Resp{Msg: err.Error()})
			return
		}
		redisUserId /* redis key */ := "user" + strconv.Itoa(int(payload.Id))
		redisToken /* redis value */, err := global.RedisCli.Get(context.Background(), redisUserId).Result()
		if err != nil {
			api.Err(ctx, api.Resp{Msg: err.Error()})
			return
		}
		if token != redisToken {
			api.Err(ctx, api.Resp{Msg: "Token error"})
			return
		}
		//! token 是否过期
		remainTTL, err := global.RedisCli.TTL(context.Background(), redisToken).Result()
		if err != nil {
			api.Err(ctx, api.Resp{Msg: err.Error()})
			return
		}
		if remainTTL < 0 {
			api.Err(ctx, api.Resp{Msg: "token expired"})
			return
		}
		//! token 续期
		expire := viper.GetDuration("redis.expire")
		var newToken string
		if remainTTL < expire/2 {
			newToken, err = util.GenToken(payload.Id, payload.Username)
			if err == nil {
				err = global.RedisCli.Set(context.Background(), redisToken, newToken, expire).Err()
			}
			if err != nil {
				api.Err(ctx, api.Resp{Msg: err.Error()})
				return
			}
			ctx.Header("Authorization", "Bearer "+newToken)
		}

		ctx.Next()
	}
}
