package main

import (
	"fmt"

	"github.com/Aazarias/ApiGomasterclass/cache"
	"github.com/Aazarias/ApiGomasterclass/db/moke"
	"github.com/Aazarias/ApiGomasterclass/service"
	"github.com/Aazarias/ApiGomasterclass/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	ListenPort string
	SecretKey  []byte
	EnvType    string
	db         struct {
		DBName string
		User   string
		Pass   string
		Port   string
	}
}

var config Config

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	config.EnvType = viper.GetString("EnvType")
	config.SecretKey = []byte(viper.GetString("SecretKey"))
	config.ListenPort = viper.GetString("ListenPort")
	// connect to DB
	config.db.DBName = viper.GetString("db.DBName")
	config.db.User = viper.GetString("db.User")
	config.db.Pass = viper.GetString("db.Pass")
	config.db.Port = viper.GetString("db.Port")
}

func main() {
	r := gin.Default()
	db := moke.New()
<<<<<<< HEAD

	secureJWT := util.MiddlJWT(config.SecretKey)
	c := cache.New()
	cacheMdw := cache.MiddlCache(c)
	s := service.New(db, c, config.SecretKey)
	r.GET("/users/:id", cacheMdw, s.GetUser)
	r.POST("/users", s.CreateUser)
	r.GET("/users", cacheMdw, s.GetAllUser)
	r.DELETE("/users/:id", secureJWT, s.DeleteUser)
	r.POST("/login", s.Login)

	r.GET("/games/:id", s.GetGames)
	r.POST("/games", s.CreateGames)
	r.GET("/games", s.GetAllGames)
	r.DELETE("/games/:id", s.DeleteGames)
	r.Run(":8081")
}
