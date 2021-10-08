package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	RunAllRoutes()
}


func RunAllRoutes() {
	r := gin.Default()

	// Исползование CORS
	//r.Use(controller.CORSMiddleware())

	// Запуск роутов
	runAllRoutes(r)

	// Запуск сервера
	_ = r.Run(":"+os.Getenv("PORT"))
}

func runAllRoutes(r *gin.Engine) {
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", PingPong)
}

// PingPong Проверка
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
