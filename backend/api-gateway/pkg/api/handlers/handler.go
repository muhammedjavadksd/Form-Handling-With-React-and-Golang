package handler

import (
	"api-gateway/pkg/models"
	"errors"
	"log"
	"net/http"

	"api-gateway/pkg/form-client/form/pb"

	"github.com/gofiber/fiber/v2"
)

type FormHandler struct {
	conn pb.FormClient
}

func NewFormHandler(conn pb.FormClient) *FormHandler {
	return &FormHandler{
		conn: conn,
	}
}

func (f *FormHandler) SaveForm(c *fiber.Ctx) error {
	var form models.FormRequest

	if err := c.BodyParser(&form); err != nil {
		log.Println(err)
		err = errors.New("could not bind form")
		c.Status(http.StatusBadGateway).SendString(err.Error())
		return err
	}

	ret, err := f.conn.InsertForm(c.Context(), &pb.FormReq{
		FullName: form.FullName,
		Email:    form.Email,
		Address:  form.Address,
		Phone:    int64(form.Phone),
	})

	if err != nil {
		c.Status(http.StatusBadGateway).SendString(err.Error())
		return err
	}
	c.Status(http.StatusAccepted).SendString(ret.String())
	return nil
}
