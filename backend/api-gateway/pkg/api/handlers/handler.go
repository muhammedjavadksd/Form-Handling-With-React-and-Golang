package handler

import (
	"api-gateway/pkg/models"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FormHandler struct {
}

func NewFormHandler() *FormHandler {
	return &FormHandler{}
}

func (f *FormHandler) SaveForm(c *fiber.Ctx) error {
	var form models.FormRequest

	if err := c.BodyParser(&form); err != nil {
		log.Println(err)
		err = errors.New("could not bind form")
		c.Status(http.StatusBadGateway).SendString(err.Error())
		return err
	}

	fmt.Println("form ", form)
	c.Status(http.StatusAccepted).SendString(form.Email)
	return nil
}
