package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	Service *service.TaskService
}

func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{Service: service}
}

func (c *TaskController) GetAllTasks(ctx echo.Context) error {
	tasks, err := c.Service.GetAllTasks()
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve tasks")
	}
	return utils.Respond(ctx, http.StatusOK, tasks, "Success")
}

func (c *TaskController) CreateTask(ctx echo.Context) error {
	var task model.Task
	if err := ctx.Bind(&task); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, err.Error())
	}
	if err := c.Service.CreateTask(task); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	return utils.Respond(ctx, http.StatusCreated, nil, "Task created successfully")
}

func (c *TaskController) GetOneTask(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid task ID")
	}
	task, err := c.Service.GetOneTask(id)

	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}

	return utils.Respond(ctx, http.StatusOK, task, "Success")
}

func (c *TaskController) UpdateTask(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid task ID")
	}

	var task model.Task
	if err := ctx.Bind(&task); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, err.Error())
	}

	if err := c.Service.UpdateTask(id, task); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}

	return utils.Respond(ctx, http.StatusOK, nil, "Task updated successfully")
}

func (c *TaskController) DeleteTask(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid task ID")
	}

	if err := c.Service.DeleteTask(id); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}

	return utils.Respond(ctx, http.StatusOK, nil, "Task deleted successfully")
}

func (c *TaskController) GetOpenLibraryData(ctx echo.Context) error {
	data, err := c.Service.GetOpenLibraryData()
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}

	return utils.Respond(ctx, http.StatusOK, data, "Successfully retrieved data from external API")
}
