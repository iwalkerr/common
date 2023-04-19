package jwtx

import (
	"context"
	"encoding/json"

	"github.com/golang-jwt/jwt"
)

const uidKey = "uid"

func GetJwtToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[uidKey] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// 获取用户ID
func GetUid(ctx context.Context) int64 {
	jsonUid, err := ctx.Value(uidKey).(json.Number).Int64()
	if err != nil {
		return 0
	}
	return jsonUid
}
