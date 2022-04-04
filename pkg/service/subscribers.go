package service

import (
	"errors"

	"github.com/Bloodstein/dns-shop-subscriber-backend/domain"
	"github.com/Bloodstein/dns-shop-subscriber-backend/pkg/repository"
)

const (
	get_microdata_url = "https://www.dns-shop.ru/product/microdata/{product_id}/"
)

type Subscribers struct {
	repo repository.Repository
}

func NewSubscribersService(repo repository.Repository) Subscribers {
	return Subscribers{
		repo: repo,
	}
}

func (this Subscribers) GetAll() []*domain.Subscribe {
	return this.repo.GetAll()
}

func (this Subscribers) GetOne(rowid int) *domain.Subscribe {
	return this.repo.GetOne(rowid)
}

func (this Subscribers) Create(request *domain.CreateNewRequest) (int, error) {
	return this.repo.Create(request)
}

func (this Subscribers) Delete(rowid int) error {
	if this.repo.GetOne(rowid) == nil {
		return errors.New("subscribe with that ID not found")
	}
	return this.repo.Delete(rowid)
}

func (this Subscribers) Monitoring(rowid int) error {
	return this.repo.Monitoring(rowid)
}

func (this Subscribers) PricesListening() {

}
