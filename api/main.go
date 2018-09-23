package main

import (
	"./database"
	"./service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"net/http"
	"os"
	"time"
)

var identityKey = "id"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	ID   int
	Role int
}

func main() {

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					"role": v.Role,

				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			fmt.Println(claims)
			return &User{
				ID: int(claims["id"].(float64)),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			login := loginVals.Username
			password := loginVals.Password

			id, role := database.User(login, password)
			if id > 0 {
				return &User{
						ID:   id,
						Role: role},
					nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.ID > 0 {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router := initRouter()
	router.POST("/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	admin := router.Group("/admin")
	admin.Use(authMiddleware.MiddlewareFunc())
	{
		admin.GET("/refresh_token", authMiddleware.RefreshHandler)

		admin.GET("/agencies", service.Agencies)
		admin.POST("/agencies", service.AddAgency)
		admin.GET("/agencies/:id", service.Agency)
		admin.PUT("/agencies/:id", service.PutAgency)

		admin.GET("/clients", service.Clients)
		admin.POST("/clients", service.AddClient)
		admin.GET("/clients/:id", service.Clients)
		admin.PUT("/clients/:id", service.PutClient)

		admin.GET("/gears", service.Gears)
		admin.POST("/gears", service.AddGear)
		admin.GET("/gears/:id", service.Gear)
		admin.PUT("/gears/:id", service.PutGear)

		admin.GET("/referrals", service.Referrals)
		admin.POST("/referrals", service.AddReferral)
		admin.GET("/referrals/:id", service.Referral)
		admin.PUT("/referrals/:id", service.PutReferral)

		admin.GET("/schedules", service.Schedules)
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:    []string{"POST", "GET", "OPTION", "PUT"},
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-Range"},
		AllowOriginFunc: original,
	}))
	return router
}

func original(origin string) bool {
	//TODO: Only http://localhost:3000
	return true
}

func SendSms(text string, phoneNumber string) {
	values := map[string]string{
		"api_key":    os.Getenv("NOTIFICATION_API_KEY"),
		"api_secret": os.Getenv("NOTIFICATION_API_SECRET"),
		"to":         phoneNumber,
		"from":       os.Getenv("NOTIFICATION_SMS_FROM"),
		"text":       text}
	jsonValue, _ := json.Marshal(values)
	http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(jsonValue))
}

func SendEmail(subject string, toName string, toEmail string, content string) {
	from := mail.NewEmail("Baby Go Round", "support@babygoround.org")
	to := mail.NewEmail(toName, toEmail)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := content
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SANDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
