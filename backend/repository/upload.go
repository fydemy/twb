package repository

import (
	"twb/model/domain"
	"twb/model/web"

	"github.com/go-redis/redis"
)

type UploadRepository interface {
	Set(*domain.Twb) error
	Get(string) (*web.TwbResponse, error)
}

type UploadRepositoryImpl struct {
	redis *redis.Client
}

func NewUploadRepo(redis *redis.Client) UploadRepository {
	return &UploadRepositoryImpl{
		redis: redis,
	}
}

func (repository *UploadRepositoryImpl) Set(twb *domain.Twb) error {
	return repository.redis.Set(twb.ID, twb.Path, twb.ExpiresIn).Err()
}

func (repository *UploadRepositoryImpl) Get(id string) (*web.TwbResponse, error) {
	path, err := repository.redis.Get(id).Result()

	if err != nil {
		return nil, err
	}

	return &web.TwbResponse{
		Path: path,
	}, nil
}
