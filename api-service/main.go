package main

import (
	"log"

	"github.com/PoowadolDev/blog-website-golang/adapter/database"
	"github.com/PoowadolDev/blog-website-golang/adapter/handler"
	"github.com/PoowadolDev/blog-website-golang/entity"
	"github.com/PoowadolDev/blog-website-golang/ports"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// Connect to the SQLite database
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Migrate the schema
	db.AutoMigrate(&entity.Posts{})
	if err != nil {
		log.Fatal("Failed to migrate database schema: ", err)
	}

	// Create a new posts repository
	postsDatabase := database.NewPostsDB(db)
	postsService := ports.NewPostsService(postsDatabase)
	postsHandler := handler.NewPostHandler(postsService)
	app := fiber.New()

	app.Get("/GetAllPosts", postsHandler.GetAll)
	app.Get("/GetPostById/:id", postsHandler.GetById)
	app.Post("/CreatePost", postsHandler.CreatePost)
	app.Put("/UpdatePost/:id", postsHandler.UpdatePost)
	app.Delete("/DeletePost/:id", postsHandler.DeletePost)

	app.Listen(":3000")
}
