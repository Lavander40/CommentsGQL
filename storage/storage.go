package storage

import (
	"CommentsGQL/models"
	"errors"
)

// собственно определённые ошибки связанные с хранилищем
var(
	ErrCommentsDisabled = errors.New("comments disabled or no such post")
)

// интерфейс хранилища используемого в приложении
type Storage interface {
	AddPost(*models.Post) error
	GetPosts() ([]*models.Post, error)
	AddComment(*models.Comment) error
}