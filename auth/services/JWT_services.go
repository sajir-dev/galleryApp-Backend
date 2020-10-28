package authservices

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

var secretkey = "secret"

// AuthCustomClaims defines the jwtclaims
type AuthCustomClaims struct {
	UserId bson.ObjectId `json:"id"  bson:"_id"`
	jwt.StandardClaims
}

// GenerateToken ...
// UserId will be in the format ObjectIdHex("5f992443d4adb48f8a1c9a7f")
func GenerateToken(UserId bson.ObjectId) string {
	claims := &AuthCustomClaims{
		UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 48).Unix(),
		},
	}

	// fmt.Println("claims: ", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// fmt.Println("token created: ", token)
	// encoded string
	t, err := token.SignedString([]byte(secretkey))
	// fmt.Println(t)
	if err != nil {
		fmt.Println("could not sign the token")
		panic(err)
	}

	// ValidateJWT(t)

	return t
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
		fmt.Printf("claims userid: %v , claims stdclaims: %v", claims.UserId, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
}
