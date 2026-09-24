package autenticacao

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CriarToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte("token_exemplo"))
}
