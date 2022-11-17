package api

import (
	v1 "github.com/samandar2605/post/api/v1"
	"github.com/samandar2605/post/config"
	"github.com/samandar2605/post/storage"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "github.com/samandar2605/post/api/docs" // for swagger
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

// @title           Swagger for blog api
// @version         1.0
// @description     This is a blog service api.
// @host      		localhost:8000
// @BasePath  		/v1
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")


	// Category
	apiV1.GET("/categories",handlerV1.GetCategoryAll)
	apiV1.GET("/categories/:id",handlerV1.GetCategory)
	apiV1.POST("/categories",handlerV1.CreateCategory)
	apiV1.PUT("/categories/:id",handlerV1.UpdateCategory)
	apiV1.DELETE("/categories/:id",handlerV1.DeleteCategory)

	// Like
	apiV1.GET("/likes/:id",handlerV1.GetLike)
	apiV1.GET("/likes",handlerV1.GetAllLike)
	apiV1.POST("/likes",handlerV1.CreateLike)
	apiV1.PUT("/likes/:id",handlerV1.UpdateLike)
	apiV1.DELETE("/likes/:id",handlerV1.DeleteLike)
	

	// User
	apiV1.GET("/users",handlerV1.GetUserAll)
	apiV1.GET("/users/:id",handlerV1.GetUser)
	apiV1.POST("/users",handlerV1.CreateUser)
	apiV1.PUT("/users/:id",handlerV1.UpdateUser)
	apiV1.DELETE("/users/:id",handlerV1.DeleteUser)

	// Comment
	apiV1.GET("/comments",handlerV1.GetAllComment)
	apiV1.GET("/comments/:id",handlerV1.GetComment)
	apiV1.POST("/comments",handlerV1.CreateComment)
	apiV1.PUT("/comments/:id",handlerV1.UpdateComment)
	apiV1.DELETE("/comments/:id",handlerV1.DeleteComment)

	// Post
	apiV1.GET("/posts",handlerV1.GetPostAll)
	apiV1.GET("/posts/:id",handlerV1.GetPost)
	apiV1.POST("/posts",handlerV1.CreatePost)
	apiV1.PUT("/posts/:id",handlerV1.UpdatePost)
	apiV1.DELETE("/posts/:id",handlerV1.DeletePost)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
