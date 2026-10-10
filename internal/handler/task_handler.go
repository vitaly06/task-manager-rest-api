package handler

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/vitaly06/task-manager-rest-api/internal/consts"
	"github.com/vitaly06/task-manager-rest-api/internal/service"
	"gorm.io/gorm"
)

type TaskHandler struct {
	taskS *service.TaskService
}

func NewTaskHandler(taskS *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskS: taskS,
	}
}

func (h *TaskHandler) Create(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var req CreateTaskRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var desc string
	if req.Description != nil {
		desc = *req.Description
	}

	var deadline time.Time
	if req.Deadline != nil {
		deadline = *req.Deadline
	}

	err := h.taskS.Create(req.Title, desc, deadline, userID)

	if err != nil {
		if errors.Is(err, service.EmptyTitleError) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Заголовок обязателен для заполнения",
			})
		}

		log.Printf("[%s] taskService.Create: %s", consts.Red("ERROR"), err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Внутренняя ошибка сервера",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Задача успешно создана",
	})
}

func (h *TaskHandler) GetTasks(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	tasks, err := h.taskS.GetTasks(userID)

	if err != nil {
		log.Printf("[%s] taskService.GetTasks: %s", consts.Red("ERROR"), err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Внутренняя ошибка сервера",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tasks": tasks,
	})
}

func (h *TaskHandler) GetTaskByID(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Парсим и проверяем переданный ID
	taskIDParam := c.Params("id")

	taskID, err := uuid.Parse(taskIDParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task id format",
		})
	}

	task, err := h.taskS.GetTaskByID(taskID, userID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Задача с данным id не найдена",
			})
		}

		log.Printf("[%s] taskService.GetTaskByID: %s", consts.Red("ERROR"), err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Внутренняя ошибка сервера",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"task": task,
	})
}

func (h *TaskHandler) UpdateTask(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Парсим и проверяем переданный ID
	taskIDParam := c.Params("id")

	taskID, err := uuid.Parse(taskIDParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task id format",
		})
	}

	var req UpdateTaskRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updated := service.UpdateTaskInput{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Deadline:    req.Deadline,
	}

	task, err := h.taskS.UpdateTask(taskID, userID, updated)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Задача с данным id не найдена",
			})
		}

		if errors.Is(err, service.EmptyTitleError) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Заголовок обязателен для заполнения",
			})
		}

		log.Printf("[%s] taskService.UpdateTask: %s", consts.Red("ERROR"), err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Внутренняя ошибка сервера",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Задача успешно обновлена",
		"task":    task,
	})
}

func (h *TaskHandler) DeleteTask(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Парсим и проверяем переданный ID
	taskIDParam := c.Params("id")

	taskID, err := uuid.Parse(taskIDParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task id format",
		})
	}

	err = h.taskS.DeleteTask(taskID, userID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Задача с данным id не найдена",
			})
		}

		log.Printf("[%s] taskService.DeleteTask: %s", consts.Red("ERROR"), err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Внутренняя ошибка сервера",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Задача успешно удалена",
	})
}
