package tasks_transport_http

import (
	"time"

	"github.com/Alv1ol/Todoapp/internal/core/domain"
)

type TaskDTOResponce struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"create_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID int        `json:"author_user_id"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponce{
	return TaskDTOResponce{
		ID: task.ID,
		Version: task.Version,
		Title: task.Title,
		Description: task.Description,
		Completed: task.Completed,
		CreatedAt: task.CreatedAt,
		CompletedAt: task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func taskDTOsFromDomains(tasks []domain.Task) []TaskDTOResponce {
	dtos := make([]TaskDTOResponce, len(tasks))

	for i, task := range tasks {
		dtos[i] = taskDTOFromDomain(task)
	}
	return dtos
}
