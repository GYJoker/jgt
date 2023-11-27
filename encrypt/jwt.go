package encrypt

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"reflect"
)

type JwtMap struct {
	jwt.Claims   `json:"-,omitempty"`
	Audience     string `json:"aud,omitempty"`
	ExpiresAt    int64  `json:"exp,omitempty"`
	Id           string `json:"jti,omitempty"`
	IssuedAt     int64  `json:"iat,omitempty"`
	RefreshAt    int64  `json:"rea,omitempty"`
	Issuer       string `json:"iss,omitempty"`
	Subject      uint64 `json:"sub,omitempty"`
	UserName     string `json:"user_name,omitempty"`
	HeaderImg    string `json:"header_img,omitempty"`
	Phone        string `json:"phone,omitempty"`
	MerchantCode string `json:"merchant_code,omitempty"`
}

func Jwt(data *JwtMap) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, data)
	tokenString, err := token.SignedString(GetRSA().PrivateKey)
	fmt.Println(err)
	return tokenString, err
}

func ParseJwt(tokenString string) (*JwtMap, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return GetRSA().PublicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}
	marshal, err := json.Marshal(token.Claims.(jwt.MapClaims))
	if err != nil {
		fmt.Println("[_checkSessionToken] " + err.Error())
		return nil, err
	}
	m := &JwtMap{}
	err = json.Unmarshal(marshal, m)
	if err != nil {
		fmt.Println("[_checkSessionToken] " + err.Error())
		return nil, err
	}

	return m, nil
}

func MarshalJwt(data interface{}) *JwtMap {
	if reflect.TypeOf(data).Kind() == reflect.TypeOf(JwtMap{}).Kind() {
		return data.(*JwtMap)
	}
	m := &JwtMap{}
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return m
	}
	err = json.Unmarshal(bytes, m)
	if err != nil {
		fmt.Println(err)
		return m
	}
	return m
}
