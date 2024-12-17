package internal

import (
	"TODO-List/internal/db"
	"TODO-List/internal/handler"
	"TODO-List/internal/repository/task"
	"TODO-List/internal/repository/user"
	"TODO-List/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func Run() {

	port := os.Getenv("PORT")
	log.Printf("server running on port ----------------------%s", port)
	database, err := db.LoadDatabase()
	defer func(database *sqlx.DB) {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}(database)
	if err != nil {
		panic(err)
	}

	userRepository := user.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	taskRepository := task.NewTaskRepository(database)
	taskService := service.NewTaskService(taskRepository, userRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/ready", func(c *gin.Context) {
		err := database.Ping()
		if err != nil {
			c.JSON(500, gin.H{"status": "unhealthy"})
			return
		}
		c.JSON(200, gin.H{"status": "ready"})
	})
	userHandler.RegisterEndpointsForUser(r)
	taskHandler.RegisterEndpointsForTasks(r)
	err = r.Run(":" + port)

	if err != nil {
		panic(err)
	}
}
