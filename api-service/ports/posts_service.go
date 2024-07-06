package ports

import (
	"errors"
	"fmt"

	"github.com/PoowadolDev/blog-website-golang/entity"
	"github.com/PoowadolDev/blog-website-golang/repository"
)

type PostsService interface {
	GetAll() ([]entity.Posts, error)
	GetById(id string) (entity.Posts, error)
	CreatePost(post entity.Posts) (entity.Posts, error)
	UpdatePost(id string, post entity.Posts) (entity.Posts, error)
	DeletePost(id string) error
}

type postsService struct {
	postsRepo repository.PostsRepo
}

func NewPostsService(postsRepo repository.PostsRepo) PostsService {
	return &postsService{
		postsRepo: postsRepo,
	}
}

func (p *postsService) GetAll() ([]entity.Posts, error) {
	return p.postsRepo.GetAllPosts()
}

func (p *postsService) GetById(id string) (entity.Posts, error) {
	return p.postsRepo.GetPostById(id)
}

func (p *postsService) CreatePost(post entity.Posts) (entity.Posts, error) {
	return p.postsRepo.CreatePost(post)
}

func (p *postsService) UpdatePost(id string, post entity.Posts) (entity.Posts, error) {
	if post, err := p.GetById(id); err != nil {
		return post, errors.New("not found this post")
	}
	return p.postsRepo.UpdatePost(id, post)
}

func (p *postsService) DeletePost(id string) error {
	if post, err := p.GetById(id); err != nil {
		fmt.Println(post)
		return errors.New("not found this post")
	}
	return p.postsRepo.DeletePost(id)
}
