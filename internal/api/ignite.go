package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/io-boxies/api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           InOut Box API Document
// @version         1.0
// @description     인아웃박스 API 문서임당
// @termsOfService  http://swagger.io/terms/

// @contact.name   최성필
// @contact.url    https://github.com/seongpil0948
// @contact.email  seongpil0948@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description                 Description for what is this security definition being used

// @securitydefinitions.oauth2.application  OAuth2Application
// @tokenUrl                                https://example.com/oauth/token
// @scope.write                             Grants write access
// @scope.admin                             Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit  OAuth2Implicit
// @authorizationUrl                     https://example.com/oauth/authorize
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.password  OAuth2Password
// @tokenUrl                             https://example.com/oauth/token
// @scope.read                           Grants read access
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode  OAuth2AccessCode
// @tokenUrl                               https://example.com/oauth/token
// @authorizationUrl                       https://example.com/oauth/authorize
// @scope.admin                            Grants read and write access to administrative information

// func main() {
// 	r := gin.Default()

// 	c := controller.NewController()

// 	v1 := r.Group("/api/v1")
// 	{
// 		accounts := v1.Group("/accounts")
// 		{
// 			accounts.GET(":id", c.ShowAccount)
// 			accounts.GET("", c.ListAccounts)
// 			accounts.POST("", c.AddAccount)
// 			accounts.DELETE(":id", c.DeleteAccount)
// 			accounts.PATCH(":id", c.UpdateAccount)
// 			accounts.POST(":id/images", c.UploadAccountImage)
// 		}
// 		bottles := v1.Group("/bottles")
// 		{
// 			bottles.GET(":id", c.ShowBottle)
// 			bottles.GET("", c.ListBottles)
// 		}
// 		admin := v1.Group("/admin")
// 		{
// 			admin.Use(auth())
// 			admin.POST("/auth", c.Auth)
// 		}
// 		examples := v1.Group("/examples")
// 		{
// 			examples.GET("ping", c.PingExample)
// 			examples.GET("calc", c.CalcExample)
// 			examples.GET("groups/:group_id/accounts/:account_id", c.PathParamsExample)
// 			examples.GET("header", c.HeaderExample)
// 			examples.GET("securities", c.SecuritiesExample)
// 			examples.GET("attribute", c.AttributeExample)
// 		}
// 	}
// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// 	r.Run(":8080")
// }

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})	
	api := r.Group("/api")
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	docs.SwaggerInfo.BasePath = "/api"
	{
		eg := api.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	// /api/swagger/index.html
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Printf("GOGOGO")
	r.Run(":8080")
}

// @BasePath  /api

// PingExample godoc
// @Summary  ping example
// @Schemes
// @Description  do ping
// @Tags         example
// @Accept       json
// @Produce      json
// @Success      200  {string}  Helloworld
// @Router       /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// func auth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if len(c.GetHeader("Authorization")) == 0 {
// 			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
// 			c.Abort()
// 		}
// 		c.Next()
// 	}
// }
