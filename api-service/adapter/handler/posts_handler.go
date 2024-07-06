package handler

import (
	"github.com/PoowadolDev/blog-website-golang/entity"
	"github.com/PoowadolDev/blog-website-golang/ports"
	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	postsService ports.PostsService
}

func NewPostHandler(postsService ports.PostsService) *PostHandler {
	return &PostHandler{postsService: postsService}
}

func (P *PostHandler) GetAll(c *fiber.Ctx) error {
	var posts []entity.Posts
	var err error
	if posts, err = P.postsService.GetAll(); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "data": posts})
}

func (p *PostHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var post entity.Posts
	var err error
	if post, err = p.postsService.GetById(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "data": post})
}

func (p *PostHandler) CreatePost(c *fiber.Ctx) error {
	var post entity.Posts
	var err error
	if err = c.BodyParser(&post); err != nil {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if post, err = p.postsService.CreatePost(post); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// p.c.Status(201)
	return c.Status(201).JSON(post)
}

func (p *PostHandler) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post entity.Posts
	var err error
	if err = c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	if post, err = p.postsService.UpdatePost(id, post); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return c.Status(200).JSON(post)
}

func (p *PostHandler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := p.postsService.DeletePost(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Post successfully deleted"})
}
