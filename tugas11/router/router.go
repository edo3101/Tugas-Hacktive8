package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tesjwt.go/controllers"
	"tesjwt.go/middlewares"
)

// @title Simple API
// @version 1.0
// @description This is a api to add photos, comments, and store the social media of users
// @termsOfService http://swagger.io/terms
// @contact.name API Support
// @contact.email redhomayan@gmail.com
// @license.name Apache 2.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @license.url http://www.apache.org/licenses/license-2.0.html
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		// Create
		userRouter.POST("/register", controllers.UserRegister)
		// Read
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		// Create
		productRouter.POST("/", controllers.CreateProduct)
		// Read
		productRouter.GET("/", controllers.FindAllProduct)
		// Update
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		// Delete
		productRouter.DELETE("/:productId", middlewares.ProductAuthorization(), controllers.DeleteProduct)
		// Read
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.FindProductById)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
