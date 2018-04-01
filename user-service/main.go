package main

import (
	"log"
	"fmt"

	pb "github.com/craigmpeters/shopper/user-service/proto/user"
	"github.com/micro/go-micro"
	_ "github.com/micro/mdns"
)

func main()  {
	// Create Database connection
	db, err := createConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Automigrate Database
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}
	tokenService := &TokenService{repo}

	// Initialise Microservice
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}


}
