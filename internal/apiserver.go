package internal

import (
	"TODO-List/internal/db"
	"TODO-List/internal/handler"
	"TODO-List/internal/logger"
	meth "TODO-List/internal/prometheus"
	"TODO-List/internal/repository/task"
	"TODO-List/internal/repository/user"
	"TODO-List/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"os"
	"time"
)

func Run() {

	port := os.Getenv("PORT")
	log.Printf("server running on port ----------------------%s", port)
	lokiURL := os.Getenv("LOKI_URL")
	logger := logger.NewLokiLogger("barebuXxX_1337", 1, lokiURL)
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
	meth.InitMetrics()
	r.GET("/metrics", func(c *gin.Context) {
		meth.RequestCounter.Inc()
		start := time.Now()
		logger.Log("Got request from prometheus", "info")
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues(c.Request.Method, "200").Observe(duration)
	})

	r.GET("/healthz", func(c *gin.Context) {
		logger.Log("checked_health", "info")
		meth.RequestCounter.Inc()
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/ready", func(c *gin.Context) {
		meth.RequestCounter.Inc()
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
