package database

import (
	"github.com/PoowadolDev/blog-website-golang/entity"
	"gorm.io/gorm"
)

type PostsDB struct {
	db *gorm.DB
}

func NewPostsDB(db *gorm.DB) *PostsDB {
	return &PostsDB{db: db}
}

func (p *PostsDB) GetAllPosts() ([]entity.Posts, error) {
	var posts []entity.Posts
	if err := p.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostsDB) GetPostById(id string) (entity.Posts, error) {
	var post entity.Posts
	if err := p.db.Where("id = ?", id).First(&post).Error; err != nil {
		return entity.Posts{}, err
	}
	return post, nil
}

func (p *PostsDB) CreatePost(post entity.Posts) (entity.Posts, error) {
	if err := p.db.Create(&post).Error; err != nil {
		return entity.Posts{}, err
	}
	return post, nil
}

func (p *PostsDB) UpdatePost(id string, post entity.Posts) (entity.Posts, error) {
	if err := p.db.Where("id = ?", id).Updates(&post).Error; err != nil {
		return entity.Posts{}, err
	}
	return post, nil
}

func (p *PostsDB) DeletePost(id string) error {
	if err := p.db.Where("id = ?", id).Delete(&entity.Posts{}).Error; err != nil {
		return err
	}
	return nil
}
