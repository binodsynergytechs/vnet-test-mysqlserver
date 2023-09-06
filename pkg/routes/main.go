package routes

import (
	"vnet-mysql/config"
	"vnet-mysql/pkg/controller"
	"vnet-mysql/pkg/middlewares"
	"vnet-mysql/pkg/repository"
	"vnet-mysql/pkg/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	router.Use(gin.Recovery())
	router.GET("/", HealthCheck)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userController.GetAllUser)
		v1.POST("/users", userController.CreateUser)
		v1.GET("/users/:id", userController.GetUserByID)
		v1.PUT("/users/:id", userController.UpdateUser)
		v1.DELETE("/users/:id", userController.DeleteUser)
	}
	router.Run(config.GetConfig().App.Port)
}

func HealthCheck(c *gin.Context) {
	resp := map[string]string{
		"status":  "OK",
		"message": "Service is running",
	}
	c.JSON(200, resp)
}
