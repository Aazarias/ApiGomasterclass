package service

import (
	"log"
	"net/http"

	"github.com/Aazarias/ApiGomasterclass/db"
	"github.com/Aazarias/ApiGomasterclass/model"
	"github.com/gin-gonic/gin"
)

//Service contain interface required
type Service struct {
	db      *db.Storage
	signKey []byte
	//cache   *cache.Cache
}

/*func New(db *db.Storage, cache *cache.Cache, signKey []byte) *Service {
	return &Service{
		db:      db,
		signKey: signKey,
		cache:   cache,
	}
}*/

//New service
func New(db *db.Storage) *Service {
	return &Service{
		db: db,
	}
}

//Get a game information by ID
func (s *Service) GetGames(c *gin.Context) {
	id := c.Param("id")
	h, err := s.db.Games.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Games": h,
	})
}

//Get all game information
func (s *Service) GetAllGames(c *gin.Context) {
	h, err := s.db.Games.GetAll()
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Games": h,
	})
}

/*func (s *Service) Update(c *gin.Context) {
	var game model.Games
	err := c.BindJSON(&game)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err upodate game": err,
		})
		return
	}

	//updated, err := s.db.Games.Update(game.ID, *db.StorageGames.Update(game, ))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err upodate game": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": updated,
	})
}
*/

//Add a game and his information
func (s *Service) CreateGames(c *gin.Context) {
	var h model.Games
	err := c.BindJSON(&h)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err creation game": err,
		})
		return
	}

	_, err = s.db.Games.Create(&h)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Games": h,
	})
}

//Delete a game
func (s *Service) DeleteGames(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error id": id,
		})
		return
	}
	err := s.db.Games.DeleteByID(id)
	if err != nil {
		log.Println("Service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"delete": id,
	})
}
