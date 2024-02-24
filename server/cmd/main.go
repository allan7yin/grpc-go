package main

import (
	"fmt"
	"log"
	"net"

	"github.com/allan7yin/grpc-go/internal/config"
	interfaces "github.com/allan7yin/grpc-go/pkg/v1"
	repo "github.com/allan7yin/grpc-go/pkg/v1/repository"
	usecase "github.com/allan7yin/grpc-go/pkg/v1/usecase"
	"gorm.io/gorm"

	dbConfig "github.com/allan7yin/grpc-go/internal/db"
	"github.com/allan7yin/grpc-go/internal/models"
	handler "github.com/allan7yin/grpc-go/pkg/v1/handler/grpc"

	"google.golang.org/grpc"
)

func main() {
	// connect to the db
	db := dbConfig.DbConnect()
	migrations(db)

	c := config.ViperConfigInit()
	srvConf := c.GetServerConfig()
	listener, err := net.Listen(srvConf.GrpcProtocol, fmt.Sprintf("%s:%s", srvConf.Host, srvConf.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	fmt.Println("Server is running on port:", srvConf.GrpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	userUseCase := initUserServer(db)
	handler.NewServer(grpcServer, userUseCase)

	// start serving to the address
	log.Fatal(grpcServer.Serve(listener))

}

func initUserServer(db *gorm.DB) interfaces.UseCaseInterface {
	userRepo := repo.New(db)
	return usecase.New(userRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
