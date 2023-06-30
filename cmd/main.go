package main

import (
	"log"
	"github.com/AvinFajarF/internal/repository"
	"github.com/AvinFajarF/internal/usecase"
	"github.com/AvinFajarF/pkg/server"
	"github.com/AvinFajarF/pkg/server/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=172.17.0.2 user=root password=root dbname=perpus port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	userRepository := repository.NewPostgresUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := http.NewUserHandler(&userUsecase)

	router := server.NewRouter(userHandler)
	router.Run()
}
