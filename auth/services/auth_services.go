package authservices

import (
	"../../domain"
	"gopkg.in/mgo.v2/bson"
)

// import (
// 	"log"
// 	"time"

// 	"../../domain"
// 	jwt "github.com/appleboy/gin-jwt"
// 	"github.com/gin-gonic/gin"
// )

// var identityKey = "id"

// type login struct {
// 	Username string `form:"username" json:"username" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

// // type User struct {
// // 	UserName  string
// // 	FirstName string
// // 	LastName  string
// // }
// var User *domain.OneUser

// func authmain() {
// 	AuthMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
// 		Realm:       "test zone",
// 		Key:         []byte("secret key"),
// 		Timeout:     time.Hour,
// 		MaxRefresh:  time.Hour,
// 		IdentityKey: identityKey,
// 		PayloadFunc: func(data interface{}) jwt.MapClaims {
// 			if v, ok := data.(*domain.OneUser); ok {
// 				return jwt.MapClaims{
// 					identityKey: v.ID,
// 				}
// 			}
// 			return jwt.MapClaims{}
// 		},

// 		Authenticator: func(c *gin.Context) (interface{}, error) {
// 			var loginVals login
// 			if err := c.ShouldBind(&loginVals); err != nil {
// 				return "", jwt.ErrMissingLoginValues
// 			}
// 			username := loginVals.Username
// 			password := loginVals.Password

// 			// if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
// 			// 	return &User{
// 			// 		UserName:  userID,
// 			// 		LastName:  "Bo-Yi",
// 			// 		FirstName: "Wu",
// 			// 	}, nil
// 			// }

// 			return nil, jwt.ErrFailedAuthentication
// 		},
// 		Authorizator: func(data interface{}, c *gin.Context) bool {
// 			if v, ok := data.(*User); ok && v.UserName == "admin" {
// 				return true
// 			}

// 			return false
// 		},
// 		Unauthorized: func(c *gin.Context, code int, message string) {
// 			c.JSON(code, gin.H{
// 				"code":    code,
// 				"message": message,
// 			})
// 		},
// 		// TokenLookup is a string in the form of "<source>:<name>" that is used
// 		// to extract token from the request.
// 		// Optional. Default value "header:Authorization".
// 		// Possible values:
// 		// - "header:<name>"
// 		// - "query:<name>"
// 		// - "cookie:<name>"
// 		// - "param:<name>"
// 		TokenLookup: "header: Authorization, query: token, cookie: jwt",
// 		// TokenLookup: "query:token",
// 		// TokenLookup: "cookie:token",

// 		// TokenHeadName is a string in the header. Default value is "Bearer"
// 		TokenHeadName: "Bearer",

// 		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
// 		TimeFunc: time.Now,
// 	})

// 	if err != nil {
// 		log.Fatal("JWT Error:" + err.Error())
// 	}

// 	// When you use jwt.New(), the function is already automatically called for checking,
// 	// which means you don't need to call it again.
// 	errInit := AuthMiddleware.MiddlewareInit()

// 	if errInit != nil {
// 		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
// 	}

// 	// r.POST("/login", authMiddleware.LoginHandler)

// 	// r.NoRoute(AuthMiddleware.MiddlewareFunc(), func(c *gin.Context) {
// 	// 	claims := jwt.ExtractClaims(c)
// 	// 	log.Printf("NoRoute claims: %#v\n", claims)
// 	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
// 	// })

// }

// // IdentityHandler: func(c *gin.Context) interface{} {
// // 	claims := jwt.ExtractClaims(c)
// // 	return &User{
// // 		ID: claims[identityKey].(string),
// // 	}
// },

// LoginService ...
// type LoginService interface {
// 	LoginUser(email string, password string)
// }

// type loginInformation struct {
// 	email    string
// 	password string
// }

// var User *domain.OneUser

// func DynamicLoginService(username string, password string) LoginService {
// 	User, err := domain.GetUser(username, password)
// 	if err != nil {
// 		return
// 	}
// }

// LoginService ...
func LoginService(username string, password string) (string, string, error) {
	User, err := domain.GetUserByCred(username, password)
	// fmt.Println("from services", User)
	if err != nil {
		return "", "", err
	}

	// Generate token
	at := GenerateToken(User.UserID, User.Username)
	rt := GenerateRefreshToken(User.UserID, User.Username)
	return at, rt, nil
}

// SignupService ...
func SignupService(username string, password string) string {
	User, err := domain.CreateUser(username, password)
	if err != nil {
		return ""
	}

	// Generate token
	token := GenerateToken(User.UserID, User.Username)
	return token
}

// UserImagesService queries all images from the db that belongs to the userid
func UserImagesService(userid bson.ObjectId) ([]domain.Image, error) {
	images, err := domain.UserItems(userid)
	if err != nil {
		return nil, err
	}
	return images, err
}

// UserImageService queries single image from the db with imageid
func UserImageService(imageid string) (*domain.Image, error) {
	image, err := domain.UserItem(imageid)
	if err != nil {
		return nil, err
	}
	return image, err
}

// UserImageCreateService creates an image with given args
func UserImageCreateService(userid string, label string, name string) (*domain.Image, error) {
	image, err := domain.UserCreateItem(userid, label, name)
	if err != nil {
		return nil, err
	}
	return image, nil
}

// DeleteImageService instantiates delete operations
func DeleteImageService(imageid string) bool {
	success := domain.UserDeleteItem(imageid)
	return success
}
