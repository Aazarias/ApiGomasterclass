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
	r.GET("/games:id", s.GetGames)
	r.POST("/games", s.CreateGames)
	r.GET("/games", s.GetAllGames)
	r.DELETE("/games/:id", s.DeleteGames)
	r.Run(":8081")
}
