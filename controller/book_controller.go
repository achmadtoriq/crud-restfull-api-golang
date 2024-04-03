package controller

import (
	"github.com/gofiber/fiber/v2"
	"go_api_module/data/response"
	"go_api_module/service"
)

type BookController struct {
	bookService service.BookService
}

func NewBookController(service service.BookService) *BookController {
	return &BookController{
		bookService: service,
	}
}

func (controller *BookController) FindAll(ctx *fiber.Ctx) error {
	bookResponse := controller.bookService.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully deleted note data !",
		Data:    bookResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
