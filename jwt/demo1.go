package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"
	"time"
)

var (
	key = []byte("mySuperSecretKeyLol")
)

type User struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
	Sex  string `json: "sex"`
}

type CustomClaims struct {
	MyUser User
	jwt.StandardClaims
}

func main() {
	//sign := encodeJWT()
	//fmt.Println(sign)

	//decodeWithClaimsJWT(sign)
	//checkTokenJWT(sign)

	sign := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNeVVzZXIiOnsiTmFtZSI6ImZh" +
		"bmdrYWltaW5nIiwiQWdlIjoxMCwiU2V4Ijoi55S3In0sImV4cCI6MTUwMDAwMDAwMDAsImlhd" +
		"CI6MTU3NjU3NjE4NSwiaXNzIjoiZmFuZ2ttIn0.K7GCwRdAqSDx3e3JI6BD_sTPu2qvMqouqtPiBhSeCbk"

	ParseToken(sign)
	//checkTokenJWT(sign)
}

func encodeJWT() string {
	myUser := User{
		Name: "fangkaiming",
		Age:  10,
		Sex:  "男",
	}

	claims := CustomClaims{
		myUser,
		jwt.StandardClaims{
			ExpiresAt: 15000000000,
			Issuer:    "fangkm",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	sign, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
	}

	return sign
}

func decodeWithClaimsJWT(token string) {
	tokenType, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		//fmt.Printf("%+v", claims)
		fmt.Println(claims.MyUser.Name)
	} else {
		log.Fatalln(err)
	}
}

func checkTokenJWT(token string) {
	parsetToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return key, nil
	})

	if err != nil {
		log.Println(err)
	}

	//fmt.Println(parsetToken.Raw, "########", parsetToken.Valid, parsetToken.Header)
	fmt.Printf("%v", parsetToken.Claims)
}

func ParseToken(token string) {
	var err error
	var headerMap map[string]interface{}
	parts := strings.Split(token, ".")
	if len(parts) < 3 {
		fmt.Println(errors.New("tokens contains invalid number of segments"))
		return
	}

	//header
	var headerByte []byte
	if headerByte, err = DecodeSegmentJWT(parts[0]); err != nil {
		if strings.HasPrefix(strings.ToLower(token), "bearer") {
			log.Fatal("tokenstring should not contain 'bearer'")
		}

		log.Fatal(err)
	}

	_ = json.Unmarshal(headerByte, &headerMap)

	//fmt.Printf("%s", headerMap["alg"])
	if method, ok := headerMap["alg"].(string); ok {
		fmt.Println(method)
	}

	fmt.Printf("%s\n", headerByte)

	//claims
	var claimsByte []byte
	claimsByte, _ = DecodeSegmentJWT(parts[1])
	fmt.Printf("%s\n", claimsByte)

	//signature
	fmt.Printf("%s\n", parts[2])
}

func DecodeSegmentJWT(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}
