package service

import (
	"github.com/Bloodstein/dns-shop-subscriber-backend/domain"
	"github.com/Bloodstein/dns-shop-subscriber-backend/pkg/repository"
)

type Service struct {
	SubscribersService SubscribersService
}

type SubscribersService interface {
	GetAll() []*domain.Subscribe
	GetOne(rowid int) *domain.Subscribe
	Create(request *domain.CreateNewRequest) (int, error)
	Delete(rowid int) error
	Monitoring(rowid int) error
	PricesListening()
}

func NewService(repo repository.Repository) Service {
	return Service{
		SubscribersService: NewSubscribersService(repo),
	}
}
