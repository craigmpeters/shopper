package main

import (
	pb "github.com/craigmpeters/shopper/item-service/proto/item"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAll() ([]*pb.Item, error)
	Get(id int32) (*pb.Item, error)
	Create(item *pb.Item) error
}

type ItemRepository struct {
	db *gorm.DB
}

func (repo *ItemRepository) GetAll() ([]*pb.Item, error) {
	var items []*pb.Item
	if err := repo.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (repo *ItemRepository) Get(id int32) (*pb.Item, error) {
	var item *pb.Item
	item.Id = id
	if err := repo.db.First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (repo *ItemRepository) Create(item *pb.Item) error {
	if err := repo.db.Create(item).Error; err != nil {
		return err
	}
	return nil
}
