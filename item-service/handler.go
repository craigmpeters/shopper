package main

import (
	_ "github.com/denisenkom/go-mssqldb"

	pb "github.com/craigmpeters/shopper/item-service/proto/item"
	"golang.org/x/net/context"
)

type service struct {
	repo Repository
}

func (srv service) Get(ctx context.Context, req *pb.Item, res *pb.Response) error {
	item, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.Item = item
	return nil
}

func (srv service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	items, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Items = items
	return nil
}

func (srv service) Create(ctx context.Context, req *pb.Item, res *pb.Response) error {
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.Item = req
	return nil
}
