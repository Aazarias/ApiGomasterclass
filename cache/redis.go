package cache

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

type Cache struct {
	cli redis.Conn
}

func New() *Cache {
	p := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	return &Cache{
		cli: p.Get(),
	}
}

func (c *Cache) Set(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = c.cli.Do("SET", key, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) SetString(key string, value string) error {
	_, err := c.cli.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	value, err := c.cli.Do("GET", key)
	if err != nil {
		return nil, err
	}

	res, ok := value.([]byte)
	if !ok {
		return nil, errors.New("cache: not able to cast")
	}
	return res, nil
}

func MiddlCache(c *Cache) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		value, err := c.Get(ctx.Request.URL.String())
		log.Println("MiddlCache", value, err)
		if len(value) != 0 && err == nil {
			log.Println("get from cache the response")
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.Header().Set("content-type", "application/json")
			_, err := ctx.Writer.Write(value)
			if err != nil {
				return
			}
			ctx.Abort()
			return
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		log.Println("set cache for next response")
		err = c.SetString(ctx.Request.URL.String(), blw.body.String())
		if err != nil {
			return
		}
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
