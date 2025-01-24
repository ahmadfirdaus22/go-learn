package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/repository"
	"go-api/router"
	"go-api/service"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	taskRepo := repository.NewTaskRepository(config.DB)
	taskService := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskService)

	authController := controller.NewAuthController(userService)
	router.InitRouter(e, userController, taskController, authController)
	e.Logger.Fatal(e.Start(":8080"))
}
