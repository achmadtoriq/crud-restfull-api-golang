package router

import (
	"github.com/gofiber/fiber/v2"
	"go_api_module/controller"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"status":  "Success",
			"message": "Welcome to Golang, Fiber, and Gorm",
		})
	})

	router.Route("/notes", func(router fiber.Router) {
		router.Post("/", noteController.Create)
		router.Get("", noteController.FindAll)
	})

	router.Route("/notes/:noteId", func(router fiber.Router) {
		router.Delete("", noteController.Delete)
		router.Get("", noteController.FindById)
		router.Patch("", noteController.Update)
	})

	return router
}
