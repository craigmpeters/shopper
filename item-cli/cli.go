package main


import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/craigmpeters/shopper/item-service/proto/item"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/context"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Item , error) {
	var item *pb.Item
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &item)
	return item, err
}

func main()  {
	cmd.Init()

	// Create new client
	client := pb.NewItemServiceClient("go.micro.srv.item", microclient.DefaultClient)

	file := defaultFilename
	var token string
	if len(os.Args) > 1 {
		file = os.Args[1]
		token = os.Args[2]
	}

	item, err :=parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	r, err := client.Create(ctx, item)
	if err != nil {
		log.Fatalf("Could not create %v", err)
	}

	log.Printf("Created: %s", r.Item.Name)

	getAll, err := client.GetAll(ctx, &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list %v", err)
	}

	for _, v  := range getAll.Items {
		log.Println(v)
	}

	
}