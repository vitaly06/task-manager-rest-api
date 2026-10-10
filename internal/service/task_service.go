package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/vitaly06/task-manager-rest-api/internal/domain"
	"github.com/vitaly06/task-manager-rest-api/internal/repository"
	"gorm.io/gorm"
)

type TaskService struct {
	taskR *repository.TaskRepository
}

func NewTaskService(taskR *repository.TaskRepository) *TaskService {
	return &TaskService{
		taskR: taskR,
	}
}

type UpdateTaskInput struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Status      *string    `json:"status"`
	Deadline    *time.Time `json:"deadline"`
}

var EmptyTitleError = errors.New("The title cannot be empty")

func (s *TaskService) Create(title, description string, deadline time.Time, userID uuid.UUID) error {
	if title == "" {
		return EmptyTitleError
	}

	task := domain.Task{
		ID:        uuid.New(),
		Title:     title,
		Status:    "new",
		CreatedAt: time.Now(),
		UserID:    userID,
	}

	if description != "" {
		task.Description = &description
	}

	if !deadline.IsZero() {
		task.Deadline = &deadline
	}

	return s.taskR.Create(&task)

}

func (s *TaskService) GetTasks(UserID uuid.UUID) ([]*domain.Task, error) {
	tasks, err := s.taskR.GetTasks(UserID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.Task{}, nil
		}

		return nil, err
	}

	return tasks, nil

}

func (s *TaskService) GetTaskByID(taskID, userID uuid.UUID) (*domain.Task, error) {
	task, err := s.taskR.GetTaskByID(taskID)

	if err != nil {
		return nil, err
	}

	if task.UserID != userID {
		return nil, gorm.ErrRecordNotFound
	}

	return task, nil
}

func (s *TaskService) UpdateTask(taskID, userID uuid.UUID, input UpdateTaskInput) (*domain.Task, error) {
	task, err := s.GetTaskByID(taskID, userID)

	if err != nil {
		return nil, err
	}

	if input.Title != nil {
		if *input.Title == "" {
			return nil, EmptyTitleError
		}
		task.Title = *input.Title
	}

	if input.Description != nil {
		task.Description = input.Description
	}

	if input.Status != nil {
		task.Status = *input.Status
	}

	if input.Deadline != nil {
		task.Deadline = input.Deadline
	}

	if err := s.taskR.UpdateTask(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) DeleteTask(taskID, userID uuid.UUID) error {
	_, err := s.GetTaskByID(taskID, userID)

	if err != nil {
		return err
	}

	return s.taskR.DeleteTask(taskID)
}
