package controller

import (
	"twb/helper"
	"twb/model/web"
	"twb/service"

	"github.com/gofiber/fiber/v2"
)

type UploadController interface {
	Post(*fiber.Ctx) error
	GetById(*fiber.Ctx) error
}

type UploadControllerImpl struct {
	uploadService service.UploadService
}

func NewUploadController(uploadService service.UploadService) UploadController {
	return &UploadControllerImpl{
		uploadService: uploadService,
	}
}

func (service *UploadControllerImpl) Post(c *fiber.Ctx) error {
	file, err := c.FormFile("frame")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": helper.ERR_NOT_FOUND,
		})
	}

	id, err := service.uploadService.Add(c, web.TwbCreateRequest{
		Frame: file,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": helper.INTERNAL_SERVICE_ERROR,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func (service *UploadControllerImpl) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	twb, err := service.uploadService.GetById(web.TwbGetRequest{
		ID: id,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": helper.INTERNAL_SERVICE_ERROR,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"path": twb.Path,
	})
}
