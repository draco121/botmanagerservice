package controllers

import (
	"github.com/draco121/botmanagerservice/core"
	"github.com/draco121/common/models"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	service core.IBotService
}

func NewControllers(service core.IBotService) Controllers {
	c := Controllers{
		service: service,
	}
	return c
}

func (s *Controllers) CreateBot(c *gin.Context) {
	var bot models.Bot
	if c.ShouldBind(&bot) != nil {
		c.JSON(400, gin.H{
			"message": "data validation error",
		})
	} else {
		res, err := s.service.CreateBot(c.Request.Context(), &bot)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}
}

func (s *Controllers) UpdateBot(c *gin.Context) {
	var bot models.Bot
	if c.ShouldBind(&bot) != nil {
		c.JSON(400, gin.H{
			"message": "data validation error",
		})
	} else {
		res, err := s.service.UpdateBot(c.Request.Context(), &bot)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}
}

func (s *Controllers) DeleteBot(c *gin.Context) {
	if botId, ok := c.GetQuery("botId"); !ok {
		c.JSON(400, gin.H{
			"message": "bot id not provided",
		})
	} else {
		res, err := s.service.DeleteBot(c.Request.Context(), botId)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}

}

func (s *Controllers) GetBot(c *gin.Context) {
	if botId, ok := c.GetQuery("botId"); !ok {
		c.JSON(400, gin.H{
			"message": "bot id not provided",
		})
	} else {
		res, err := s.service.DeleteBot(c.Request.Context(), botId)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}
	if projectId, ok := c.GetQuery("projectId"); !ok {
		c.JSON(400, gin.H{
			"message": "bot id not provided",
		})
	} else {
		res, err := s.service.DeleteBot(c.Request.Context(), projectId)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}
}
