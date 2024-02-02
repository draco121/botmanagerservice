package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/draco121/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IBotRepository interface {
	InsertOne(ctx context.Context, bot *models.Bot) (*models.Bot, error)
	UpdateOne(ctx context.Context, bot *models.Bot) (*models.Bot, error)
	FindOneById(ctx context.Context, id string) (*models.Bot, error)
	FindOneByName(ctx context.Context, name string) (*models.Bot, error)
	DeleteOneById(ctx context.Context, id string) (*models.Bot, error)
	FindManyByProjectId(ctx context.Context, projectId string) (*[]models.Bot, error)
}

type botRepository struct {
	IBotRepository
	db *mongo.Database
}

func NewBotRepository(database *mongo.Database) IBotRepository {
	repo := botRepository{db: database}
	return &repo
}

func (ur botRepository) InsertOne(ctx context.Context, bot *models.Bot) (*models.Bot, error) {
	result, _ := ur.FindOneByName(ctx, bot.Name)
	if result != nil {
		return nil, fmt.Errorf("record exists")
	} else {
		bot.ID = primitive.NewObjectID()
		_, err := ur.db.Collection("bots").InsertOne(ctx, bot)
		if err != nil {
			return nil, err
		}
	}
	return bot, nil
}

func (ur botRepository) UpdateOne(ctx context.Context, bot *models.Bot) (*models.Bot, error) {
	filter := bson.M{"_id": bot.ID}
	update := bson.M{"$set": bson.M{
		"apiKey": bot.ApiKey,
	}}
	result := models.Bot{}
	err := ur.db.Collection("bots").FindOneAndUpdate(ctx, filter, update).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur botRepository) FindOneById(ctx context.Context, id string) (*models.Bot, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	} else {
		filter := bson.D{{Key: "_id", Value: objectId}}
		result := models.Bot{}
		err := ur.db.Collection("bots").FindOne(ctx, filter).Decode(&result)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		} else {
			return &result, nil
		}
	}
}

func (ur botRepository) FindOneByName(ctx context.Context, name string) (*models.Bot, error) {
	filter := bson.D{{Key: "email", Value: name}}
	result := models.Bot{}
	err := ur.db.Collection("bots").FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur botRepository) DeleteOneById(ctx context.Context, id string) (*models.Bot, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	} else {
		filter := bson.D{{Key: "_id", Value: objectId}}
		result := models.Bot{}
		err := ur.db.Collection("bots").FindOneAndDelete(ctx, filter).Decode(&result)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		} else {
			return &result, nil
		}
	}
}

func (ur botRepository) FindManyByProjectId(ctx context.Context, projectId string) (*[]models.Bot, error) {
	var bots []models.Bot
	filter := bson.D{{Key: "projectId", Value: projectId}}
	cur, err := ur.db.Collection("bots").Find(ctx, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		if err = cur.All(context.TODO(), &bots); err != nil {
			return nil, err
		}
		return &bots, nil
	}
}
