package service

import (
	"log"
	"net/http"

	"github.com/Aazarias/ApiGomasterclass/model"
	"github.com/Aazarias/ApiGomasterclass/util"
	"github.com/gin-gonic/gin"
)

// Get a user by his ID
func (s *Service) GetUser(c *gin.Context) {
	id := c.Param("id")
	u, err := s.db.User.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

//show all user
func (s *Service) GetAllUser(c *gin.Context) {
	us, err := s.db.User.GetAll()
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": us,
	})
}

//add a user, ID must be unique
func (s *Service) CreateUser(c *gin.Context) {
	var u model.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	if u.Password == nil || len(*u.Password) == 0 || len(u.Email) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "need email and password",
		})
		return
	}

	_, err = s.db.User.Create(&u)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

//Delete a user
func (s *Service) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error id": id,
		})
		return
	}
	err := s.db.User.DeleteByID(id)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"delete": id,
	})
}

//Check if email and password refer to a user
func (s *Service) Login(c *gin.Context) {

	var l model.LoginUser
	err := c.BindJSON(&l)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	if l.Password == nil || len(*l.Password) == 0 || len(l.Email) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "not authorized",
		})
		return
	}

	u, err := s.db.User.GetByEmail(l.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"email": l.Email,
		})
		return
	}

	if *u.Password != *l.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "not authorized",
		})
		return
	}

	jwtVal, err := util.CreateJWT(s.signKey, u.Id, u.Firstname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt": jwtVal,
	})
}
