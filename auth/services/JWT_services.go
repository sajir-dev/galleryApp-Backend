package authservices

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

var secretkey = "secret"
var refreshsecret = "secret"

// AuthCustomClaims defines the jwtclaims
type AuthCustomClaims struct {
	UserID   bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	jwt.StandardClaims
}

// GenerateToken ...
// UserId will be in the format ObjectIdHex("5f992443d4adb48f8a1c9a7f")
func GenerateToken(userID bson.ObjectId, username string) string {
	claims := &AuthCustomClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
	}

	// fmt.Println("claims: ", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// fmt.Println("token created: ", token)
	// encoded string
	at, err := token.SignedString([]byte(secretkey))
	// fmt.Println(t)
	if err != nil {
		// fmt.Println("could not sign the token")
		panic(err)
	}

	// ValidateJWT(t)

	return at
}

// GenerateRefreshToken ...
func GenerateRefreshToken(UserID bson.ObjectId, username string) string {
	claims := &AuthCustomClaims{
		UserID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	rt, err := token.SignedString([]byte(refreshsecret))
	if err != nil {
		fmt.Println("Could not sign the refresh token")
		panic(err)
	}

	return rt
}

// ValidateJWT helps to validate the token
func ValidateJWT(jwtFromHeader string) {
	// fmt.Println(jwtFromHeader)
	token, err := jwt.ParseWithClaims(
		jwtFromHeader,
		&AuthCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretkey), nil
		})

	if claims, ok := token.Claims.(*AuthCustomClaims); ok && token.Valid {
		// fmt.Println(claims)
		// fmt.Println(claims.UserId)
		// code, _ := hex.DecodeString(claims.UserId)
		fmt.Printf("claims userid: %v , claims username: %v claims stdclaims: %v", claims.UserID, claims.Username, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
}

// InvalidateJWT ...
func InvalidateJWT(jwtFromHeader string) {
	token, err := jwt.ParseWithClaims(
		jwtFromHeader,
		&AuthCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretkey), nil
		},
	)
	if err != nil {
		fmt.Println("could not parse token with claims")
		return
	}
	if claims, ok := token.Claims.(AuthCustomClaims); ok && token.Valid {
		claims.StandardClaims.ExpiresAt = time.Now().Unix()
		fmt.Printf("claims userid: %v , claims username: %v claims stdclaims: %v", claims.UserID, claims.Username, claims.StandardClaims.ExpiresAt)
		return
	}
	return
}

// RefreshHandler ...
func RefreshHandler(rt string) (string, string, error) {
	token, err := jwt.ParseWithClaims(
		rt,
		&AuthCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(refreshsecret), nil
		},
	)
	if err != nil {
		return "", "", errors.New("Invalid token")
	}
	if claims, ok := token.Claims.(*AuthCustomClaims); ok && token.Valid {
		userid := claims.UserID
		username := claims.Username
		at := GenerateToken(userid, username)
		rt := GenerateRefreshToken(userid, username)
		return at, rt, nil
	}
	return "", "", errors.New("Unable to login")
}
