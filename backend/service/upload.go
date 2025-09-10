package service

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"twb/helper"
	"twb/model/domain"
	"twb/model/web"
	"twb/repository"

	"github.com/gofiber/fiber/v2"
)

type UploadService interface {
	Add(*fiber.Ctx, web.TwbCreateRequest) (string, error)
	GetById(web.TwbGetRequest) (*web.TwbResponse, error)
}

type UploadServiceImpl struct {
	uploadRepository repository.UploadRepository
}

func NewUploadService(uploadRepository repository.UploadRepository) UploadService {
	return &UploadServiceImpl{
		uploadRepository: uploadRepository,
	}
}

func (service *UploadServiceImpl) Add(c *fiber.Ctx, req web.TwbCreateRequest) (string, error) {
	id, err := helper.NewBase64URL()

	if err != nil {
		return "", err
	}

	folder := "uploads"
	path := fmt.Sprintf("%s/%s%s", folder, id, filepath.Ext(req.Frame.Filename))

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return "", err
	}

	if err := c.SaveFile(req.Frame, path); err != nil {
		return "", err
	}

	if err := service.uploadRepository.Set(&domain.Twb{
		ID:        id,
		Path:      os.Getenv("APP_URL") + "/" + path,
		ExpiresIn: 7 * 24 * time.Hour, // staticly typed
	}); err != nil {
		return "", err
	}

	return id, nil
}

func (service *UploadServiceImpl) GetById(req web.TwbGetRequest) (*web.TwbResponse, error) {
	twb, err := service.uploadRepository.Get(req.ID)

	if err != nil {
		return nil, err
	}

	return twb, nil
}
