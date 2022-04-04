package repository

import "github.com/Bloodstein/dns-shop-subscriber-backend/domain"

type Repository interface {
	GetAll() []*domain.Subscribe
	GetOne(rowid int) *domain.Subscribe
	Create(request *domain.CreateNewRequest) (int, error)
	Delete(rowid int) error
	Monitoring(rowid int) error
}

func NewRepository() Repository {
	return JsonRepository{}
}
