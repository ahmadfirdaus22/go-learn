package router

import (
	"go-api/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, userController *controller.UserController, taskController *controller.TaskController, authController *controller.AuthController) {

	userGroup := e.Group("/users")

	userGroup.GET("", userController.GetAllUsers)
	userGroup.GET("/:id", userController.GetOneUser)
	userGroup.POST("", userController.CreateUser)
	userGroup.PUT("/:id", userController.UpdateUser)
	userGroup.DELETE("/:id", userController.DeleteUser)
	userGroup.GET("/openlibrary", userController.GetOpenLibraryData)

	taskGroup := e.Group("/task")
	taskGroup.GET("/tasks", taskController.GetAllTasks)
	taskGroup.GET("/tasks/:id", taskController.GetOneTask)
	taskGroup.POST("/tasks", taskController.CreateTask)
	taskGroup.PUT("/tasks/:id", taskController.UpdateTask)
	taskGroup.DELETE("/tasks/:id", taskController.DeleteTask)

	e.POST("/login", authController.LoginHandler)
	e.POST("/logout", authController.LoginHandler)
}
