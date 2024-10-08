package response

import (
	"time"
	"todo-app--go-gin/domain"
)

type TodoResponse struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTodoResponse(todo domain.Todo) TodoResponse {
	return TodoResponse{
		Id:          todo.Id,
		UserId:      todo.UserId,
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
