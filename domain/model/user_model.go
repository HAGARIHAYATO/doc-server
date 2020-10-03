package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/k-washi/jwt-decode/jwtdecode"
	"log"
	"time"
)

type User struct {
	ID int64 `db:"id"`
	Name string `db:"name"`
	Email string `db:"email"`
	Password string `db:"password"`
	Docs []*Doc
	//Bundles []*Bundle
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func (u User)CreateToken() (string, error) {
	var err error

	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"iss":   u.Name,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func DecodeToken(token string) (string, error) {
	hCS, err := jwtdecode.JwtDecode.DecomposeFB(token)
	if err != nil {
		return "", err
	}
	payload, err := jwtdecode.JwtDecode.DecodeClaimFB(hCS[1])
	if err != nil {
		return "", err
	}
	return payload.Email, nil
}
