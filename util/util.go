/*
 * @Auther: Biny
 * @Description: jwt token工具类,其他公共工具类
 * @Date: 2022-01-19 19:04:25
 * @LastEditTime: 2022-01-23 12:11:29
 */
package util

import (
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT : HEADER PAYLOAD SIGNATURE
const (
	SecretKEY              string = "JWT-Secret-Key"
	DEFAULT_EXPIRE_SECONDS int    = 180 // default expired 1 minute
	PasswordHashBytes             = 16
)

// This struct is the payload
type MyCustomClaims struct {
	UserID int64 `json:"UserID"`
	jwt.StandardClaims
}

// This struct is the parsing of token payload
type JwtPayload struct {
	Username  string `json:"Username"`
	UserID    int64  `json:"UserID"`
	IssuedAt  int64  `json:"Iat"`
	ExpiresAt int64  `json:"Exp"`
}

//generate token
func GenerateToken(userID int64, userName string, expiredSeconds int) (tokenString string, err error) {

	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}

	// Create the Claims
	mySigningKey := []byte(SecretKEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	// log.Println("userID ", userID, " Token will be expired at ", time.Unix(expireAt, 0))

	claims := MyCustomClaims{
		userID,
		jwt.StandardClaims{
			Issuer:    userName,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireAt,
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Signs the token with a secret
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", errors.New("error: failed to generate token")
	}

	return tokenStr, nil
}

//validate token
func ValidateToken(tokenString string) (*JwtPayload, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface {
		}, error) {

			return []byte(SecretKEY), nil
		})

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {

		// log.Println("ok && token valid")
		// log.Println(claims.UserID, claims.StandardClaims.ExpiresAt)
		// log.Println("Token was issued at ", time.Now().Unix())
		// log.Println("Token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))

		return &JwtPayload{
			Username:  claims.StandardClaims.Issuer,
			UserID:    claims.UserID,
			IssuedAt:  claims.StandardClaims.IssuedAt,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
		}, nil
	} else {

		fmt.Println(err)
		return nil, errors.New("error: failed to validate token")
	}
}

//update token
func RefreshToken(tokenString string) (newTokenString string, err error) {

	// get previous token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface {
		}, error) {

			return []byte(SecretKEY), nil
		})

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {

		return "", err
	}

	mySigningKey := []byte(SecretKEY)
	expireAt := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix() //new expired
	newClaims := MyCustomClaims{

		claims.UserID,
		jwt.StandardClaims{

			Issuer:    claims.StandardClaims.Issuer, //name of token issue
			IssuedAt:  time.Now().Unix(),            //time of token issue
			ExpiresAt: expireAt,                     // new expired
		},
	}

	// generate new token with new claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	// sign the token with a secret
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {

		return "", errors.New("error: failed to generate new fresh json web token")
	}

	return tokenStr, nil
}

func Password(value string) (v string) {
	o1 := sha1.New()
	o1.Write([]byte("Biny_"))
	has1 := o1.Sum(nil)
	v1 := fmt.Sprintf("%x", has1)

	data2 := []byte(value)
	has2 := md5.Sum(data2)
	v2 := fmt.Sprintf("%x", has2)

	encrypt := []byte("_encrypt")
	has3 := md5.Sum(encrypt)
	v3 := fmt.Sprintf("%x", has3)

	o2 := sha1.New()
	o2.Write([]byte(value))
	has4 := o2.Sum(nil)
	v4 := fmt.Sprintf("%x", has4)

	o5 := sha1.New()
	o5.Write([]byte(v1 + v2 + v3 + v4))
	has5 := o5.Sum(nil)
	v5 := fmt.Sprintf("%x", has5)

	return v5
}
