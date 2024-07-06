package repository

import (
	"github.com/PoowadolDev/blog-website-golang/entity"
)

type PostsRepo interface {
	GetAllPosts() ([]entity.Posts, error)
	GetPostById(id string) (entity.Posts, error)
	CreatePost(post entity.Posts) (entity.Posts, error)
	UpdatePost(id string, post entity.Posts) (entity.Posts, error)
	DeletePost(id string) error
}
