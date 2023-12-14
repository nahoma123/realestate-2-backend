package initiator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/handler/middleware"
	"visitor_management/platform/logger"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"

	// "github.com/swaggo/gin-swagger"s

	"go.uber.org/zap"
)

// gin-swagger middleware
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)

		fmt.Printf("%s %s %s %s\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			latency,
		)
	}
}

func ResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		c.Next()

		fmt.Printf("%d %s %s\n",
			c.Writer.Status(),
			c.Request.Method,
			c.Request.RequestURI,
		)
	}
}

func Initiate() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(context.Background(), "Error loading .env file")
	// }
	london, err := time.LoadLocation("Europe/London")
	if err != nil {
		// Handle the error if the time zone cannot be loaded
		panic(err)
	}

	// Set the default time zone to London
	time.Local = london

	log := logger.New(InitLogger())
	log.Info(context.Background(), "logger initialized")

	log.Info(context.Background(), "initializing config")
	configName := "config"
	if name := os.Getenv("CONFIG_NAME"); name != "" {
		configName = name
		log.Info(context.Background(), fmt.Sprintf("config name is set to %s", configName))
	} else {
		log.Info(context.Background(), "using default config name 'config'")
	}
	log.Info(context.Background(), "config initialized")

	log.Info(context.Background(), "initializing database")

	log.Info(context.Background(), "database initialized")

	log.Info(context.Background(), "initializing migration")
	log.Info(context.Background(), "migration initialized")

	log.Info(context.Background(), "initializing persistence layer")

	db, err := createDB(constant.GetConfig().DB_HOST, constant.GetConfig().DB_USER,
		constant.GetConfig().DB_PASSWORD, constant.GetConfig().DATABASE_NAME,
		constant.GetConfig().PORT)
	if err != nil {
		panic(err)
	}
	InitiateMigration(db, log)

	CreateIndexes(log, db)

	persistence := InitPersistence(db, log)
	Cron(persistence)
	log.Info(context.Background(), "persistence layer initialized")

	log.Info(context.Background(), "initializing platform layer")
	platformLayer := InitPlatformLayer(log)
	log.Info(context.Background(), "platform layer initialized")

	log.Info(context.Background(), "initializing module")
	module := InitModule(persistence, "", platformLayer, log)
	log.Info(context.Background(), "module initialized")

	log.Info(context.Background(), "initializing handler")
	handler := InitHandler(module, log)
	log.Info(context.Background(), "handler initialized")

	log.Info(context.Background(), "initializing server")
	server := gin.Default()
	// server.Use(middleware.GinLogger(log))
	server.Use(ginzap.RecoveryWithZap(log.GetZapLogger().Named("gin.recovery"), true))
	server.Use(middleware.ErrorHandler())
	if true {
		server.Use(InitCORS())
	}
	server.Use(RequestLogger())
	server.Use(ResponseLogger())

	log.Info(context.Background(), "server initialized")

	log.Info(context.Background(), "initializing metrics route")
	InitMetricsRoute(server, log)
	log.Info(context.Background(), "metrics route initialized")

	log.Info(context.Background(), "initializing router")
	v1 := server.Group("")
	InitRouter(server, v1, handler, module, log, "")
	log.Info(context.Background(), "router initialized")

	srv := &http.Server{
		Addr:    ":" + constant.GetConfig().SERVER_PORT,
		Handler: server,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	port, _ := strconv.Atoi(constant.GetConfig().SERVER_PORT)
	go func() {
		log.Info(context.Background(), "server started",
			zap.String("host", constant.GetConfig().SERVER_PORT),
			zap.Int("port", port))
		log.Info(context.Background(), fmt.Sprintf("server stopped with error %v", srv.ListenAndServe()))
	}()
	sig := <-quit
	log.Info(context.Background(), fmt.Sprintf("server shutting down with signal %v", sig))
	timeout, _ := strconv.Atoi(constant.GetConfig().SERVER_TIMEOUT)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout))
	defer cancel()

	log.Info(ctx, "shutting down server")
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("error while shutting down server: %v", err))
	} else {
		log.Info(context.Background(), "server shutdown complete")
	}
}

func Cron(pr Persistence) {
	c := cron.New()

	// Add a cron job that runs at midnight every day
	c.AddFunc("0 0 0 * * *", func() {
	})

	// Start the cron job
	c.Start()
}
