package server

import (
	"context"
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"microblog-api/auth"
	http2 "microblog-api/auth/delivery/http"
	"microblog-api/auth/repositories"
	"microblog-api/auth/services"
	"microblog-api/post"
	http4 "microblog-api/post/delivery/http"
	repositories3 "microblog-api/post/repositories"
	services3 "microblog-api/post/services"
	"microblog-api/profile"
	http3 "microblog-api/profile/delivery/http"
	repositories2 "microblog-api/profile/repositories"
	services2 "microblog-api/profile/services"
	minio2 "microblog-api/storage/minio"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	authService    auth.Service
	profileService profile.Service
	postService    post.Service
	server         *http.Server
}

func NewApp() *App {
	db := initDb()
	endpoint := os.Getenv("s3_conn")
	client, _ := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("s3_id"), os.Getenv("s3_secret"), ""),
		Secure: false,
	})
	storage := minio2.NewMinioStorage(client, "app")
	authRepo := repositories.NewPostgresRepository(db)
	profileRepo := repositories2.NewPostgresRepository(db)
	postRepo := repositories3.NewPostgresRepository(db)
	profileService := services2.NewProfileService(profileRepo, storage)
	return &App{
		authService:    services.NewUserService(authRepo, profileService, os.Getenv("password_salt"), os.Getenv("signing_key"), 1000000),
		profileService: profileService,
		postService:    services3.NewPostService(postRepo, storage),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		//cors.Default(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowHeaders:     []string{"Access-Control-Allow-Origin", "Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Cache-Control", "X-Requested-With"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
	)
	m := http2.NewAuthMiddleware(a.authService)
	apiGroup := router.Group("/api")
	http2.RegisterHTTPEndpoints(router, a.authService, m)
	http3.RegisterHTTPEndpoints(apiGroup, a.profileService, a.postService, m)
	http4.RegisterHTTPEndpoints(apiGroup, a.postService, m)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	a.server = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Fatalf("Server error: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return a.server.Shutdown(ctx)
}

func initDb() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("db_conn"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return db
}
