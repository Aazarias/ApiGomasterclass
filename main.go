package main

import (
	"github.com/Aazarias/ApiGomasterclass/db/moke"
	"github.com/Aazarias/ApiGomasterclass/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := moke.New()
	s := service.New(db)
	r.GET("/heroes/:id", s.GetGames)
	r.POST("/heroes", s.CreateGames)
	r.GET("/heroes", s.GetAllGames)
	r.DELETE("/heroes/:id", s.DeleteGames)
	r.Run(":8081")
}
