package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mirasildev/exam_project_2.0/api/v1"
	"github.com/mirasildev/exam_project_2.0/config"
	"github.com/mirasildev/exam_project_2.0/storage"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/mirasildev/exam_project_2.0/api/docs" // for swagger
)

type RouterOptions struct {
	Cfg      *config.Config
	Storage  storage.StorageI
	InMemory storage.InMemoryStorageI
}

// / @title           Swagger for booking api
// @version         1.0
// @description     This is a booking service api.
// @host      localhost:8000
// @BasePath  /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:      opt.Cfg,
		Storage:  opt.Storage,
		InMemory: opt.InMemory,
	})

	router.Static("/media", "./media")

	apiV1 := router.Group("/v1")

	apiV1.POST("/users/", handlerV1.AuthMiddleware, handlerV1.CreateUser)
	apiV1.GET("/users/:id", handlerV1.GetUser)
	apiV1.GET("/users/", handlerV1.GetAllUsers)
	apiV1.PUT("/user/:id", handlerV1.UpdateUser)
	apiV1.DELETE("/user/:id", handlerV1.DeleteUser)

	apiV1.POST("/hotels/", handlerV1.AuthMiddleware, handlerV1.CreateHotel)
	apiV1.GET("/hotels/:id", handlerV1.GetHotel)
	apiV1.GET("/hotels/", handlerV1.GetAllHotels)
	apiV1.PUT("/hotels/:id", handlerV1.UpdateHotel)
	apiV1.DELETE("/hotels/:id", handlerV1.DeleteHotel)

	apiV1.POST("/rooms/", handlerV1.AuthMiddleware, handlerV1.CreateRoom)
	apiV1.GET("/rooms/:id", handlerV1.GetRoom)
	apiV1.GET("/rooms/", handlerV1.GetAllRooms)
	apiV1.PUT("/rooms/:id", handlerV1.UpdateRoom)
	apiV1.DELETE("/rooms/:id", handlerV1.DeleteRoom)

	apiV1.POST("/bookings/", handlerV1.AuthMiddleware, handlerV1.CreateBooking)
	apiV1.GET("/bookings/:id", handlerV1.GetBooking)
	apiV1.GET("/bookings/", handlerV1.GetAllBookings)
	apiV1.PUT("/bookings/:id", handlerV1.UpdateBooking)
	apiV1.DELETE("/bookings/:id", handlerV1.DeleteBooking)

	apiV1.POST("/auth/register", handlerV1.Register)
	apiV1.POST("/auth/register-as-a-partner", handlerV1.RegisterAsAPartner)
	apiV1.POST("/auth/login", handlerV1.Login)
	apiV1.POST("/auth/verify", handlerV1.Verify)
	apiV1.POST("/auth/forgot-password", handlerV1.ForgotPassword)
	apiV1.POST("/auth/verify-forgot-password", handlerV1.VerifyForgotPassword)
	apiV1.POST("/auth/update-password", handlerV1.AuthMiddleware, handlerV1.UpdatePassword)

	apiV1.POST("/file-upload", handlerV1.UploadFile)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router

}