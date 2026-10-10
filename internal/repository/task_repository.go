package repository

import (
	"github.com/google/uuid"
	"github.com/vitaly06/task-manager-rest-api/internal/domain"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) GetTasks(userID uuid.UUID) ([]*domain.Task, error) {
	var tasks []*domain.Task

	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(taskID uuid.UUID) (*domain.Task, error) {
	var task domain.Task

	if err := r.db.Where("id = ?", taskID).First(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) UpdateTask(task *domain.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) DeleteTask(taskID uuid.UUID) error {
	return r.db.Where("id = ?", taskID).Delete(&domain.Task{}).Error
}
