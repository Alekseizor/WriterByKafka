package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var k = 0

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "127.0.0.1:29092")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (a *Application) StartServer() {
	log.Println("Server start up")
	r := gin.New()
	r.Use(CORSMiddleware())
	r.GET("/delivery", a.Delivery)
	//	r.LoadHTMLGlob("templates/*")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Println("Server down")
}

type EventDev struct {
	ID     int
	Status string
}

func (a *Application) Delivery(c *gin.Context) {
	status := [...]string{
		"Оформлен", "Сборка", "Доставка", "Готово",
	}
	var event EventDev
	event.Status = status[k%4]
	event.ID = 236322856
	c.JSON(http.StatusOK, event)
	k += 1
}
