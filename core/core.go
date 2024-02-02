package core

import (
	"context"
	"github.com/draco121/botmanagerservice/repository"
	"github.com/draco121/common/models"
)

type IBotService interface {
	CreateBot(ctx context.Context, user *models.Bot) (*models.Bot, error)
	UpdateBot(ctx context.Context, user *models.Bot) (*models.Bot, error)
	DeleteBot(ctx context.Context, id string) (*models.Bot, error)
	GetBotByName(ctx context.Context, name string) (*models.Bot, error)
	GetBotById(ctx context.Context, id string) (*models.Bot, error)
	GetBotsByProjectId(ctx context.Context, projectId string) (*[]models.Bot, error)
}

type botService struct {
	IBotService
	repo repository.IBotRepository
}

func NewBotService(repository repository.IBotRepository) IBotService {
	us := botService{
		repo: repository,
	}
	return us
}

func (s botService) CreateBot(ctx context.Context, bot *models.Bot) (*models.Bot, error) {
	return s.repo.InsertOne(ctx, bot)
}

func (s botService) UpdateBot(ctx context.Context, bot *models.Bot) (*models.Bot, error) {
	return s.repo.UpdateOne(ctx, bot)
}

func (s botService) DeleteBot(ctx context.Context, id string) (*models.Bot, error) {
	return s.repo.DeleteOneById(ctx, id)
}

func (s botService) GetBotByName(ctx context.Context, name string) (*models.Bot, error) {
	return s.repo.FindOneByName(ctx, name)
}

func (s botService) GetBotById(ctx context.Context, id string) (*models.Bot, error) {
	return s.repo.FindOneById(ctx, id)
}

func (s botService) GetBotsByProjectId(ctx context.Context, projectId string) (*[]models.Bot, error) {
	return s.repo.FindManyByProjectId(ctx, projectId)
}
