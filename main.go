package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go_api_module/config"
	"go_api_module/controller"
	"go_api_module/model"
	"go_api_module/repository"
	"go_api_module/router"
	"go_api_module/service"
	"log"
)

func main() {
	fmt.Println("Run Service...")
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load env", err)
	}

	// Init Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()
	db.Table("notes").AutoMigrate(&model.Note{})
	// init Repository
	noteRepository := repository.NewNoteRepositoryImpl(db)
	// Init Service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)
	// Init Controller
	noteController := controller.NewNoteController(noteService)
	// Init Router
	routes := router.NewRouter(noteController)

	app := fiber.New()
	app.Mount("/api", routes)
	log.Fatal(app.Listen(":3000"))
}
