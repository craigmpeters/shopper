package main

import (
	"fmt"
	pb "github.com/craigmpeters/shopper/item-service/proto/item"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	db, err := createConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.Item{})

	repo := &ItemRepository{db}

	srv := micro.NewService(
		micro.Name("go.micro.srv.item"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterItemServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
