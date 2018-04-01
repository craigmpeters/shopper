package main

import (
	"log"
	"os"

	pb "github.com/craigmpeters/shopper/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main()  {
	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name: "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name: "email",
				Usage: "Your email address",
			},
			cli.StringFlag{
				Name: "password",
				Usage: "Your Password",
			},
			cli.StringFlag{
				Name: "company",
				Usage: "Your Company",
			},

		),
	)

	service.Init(

		micro.Action(func(c *cli.Context) {

			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			r, err := client.Create(context.TODO(), &pb.User{
				Name: name,
				Email: email,
				Password: password,
				Company: company,
			})
			if err != nil {
				log.Fatalf("Could not create %v", err)
			}
			log.Printf("Created %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(),&pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users %v", err)
			}

			for _ , v := range getAll.Users {
				log.Println(v)
			}

			authResponse, err := client.Auth(context.TODO(), &pb.User{
				Email: email,
				Password: password,
			})

			if err != nil {
				log.Fatalf("Could not authenticate user %s, error: %v \n", email, err)
			}

			log.Printf("Your access token is %s \n", authResponse.Token)

			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Println(err)
	}

}
