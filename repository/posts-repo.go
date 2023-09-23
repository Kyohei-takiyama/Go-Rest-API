package repository

import (
	"github.com/Kyohei-takiyama/GoRestApi/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post , error)
	FindAll() ([]entity.Post , error)
}