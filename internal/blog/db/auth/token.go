package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/PoulDev/lgBlog/internal/blog/config"
)

func GenToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(config.JWTSecret);
	if (err != nil) {
		return "", err
	}

	return tokenString, nil;
}


func CheckToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return config.JWTSecret, nil
    })

	if (err != nil) { // parsing fallito
		return jwt.MapClaims{}, err
	}

	if (!token.Valid) {
		return jwt.MapClaims{}, errors.New("invalid token")
	}

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp := claims["exp"].(float64)

		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, errors.New("token has expired")
		}

        return claims, nil
    }

    return jwt.MapClaims{}, errors.New("failed to get claims(??)");
}
