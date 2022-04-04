package repository

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Bloodstein/dns-shop-subscriber-backend/domain"
)

const (
	subscribes_db_file_path = "database/subscribes.json"
)

type JsonRepository struct{}

func (this JsonRepository) GetAll() []*domain.Subscribe {
	return this.readAll()
}

func (this JsonRepository) GetOne(rowid int) *domain.Subscribe {
	return this.readOne(rowid)
}

func (this JsonRepository) Create(request *domain.CreateNewRequest) (int, error) {
	all := this.readAll()
	nextId := len(all) + 1
	all = append(all, &domain.Subscribe{
		ID:     nextId,
		ShopID: request.ProductId,
		Url:    request.Link,
		Name:   request.Name,
	})

	err := this.write(all)

	return nextId, err
}

func (this JsonRepository) Delete(rowid int) error {
	all := this.readAll()
	copy := make([]*domain.Subscribe, len(all)-1)

	for _, sub := range all {
		if sub.ID != rowid {
			copy = append(copy, sub)
		}
	}

	return this.write(copy)
}

func (this JsonRepository) Monitoring(rowid int) error {
	return nil
}

func (this JsonRepository) readAll() []*domain.Subscribe {
	buff, err := os.ReadFile(subscribes_db_file_path)
	if err != nil {
		log.Fatalf("fail to read subscribes.json. %s", err.Error())
	}

	var rows []*domain.Subscribe

	err = json.Unmarshal(buff, &rows)
	if err != nil {
		log.Fatalf("fail to marshal data from subscribes.json. %s", err.Error())
	}

	return rows
}

func (this JsonRepository) readOne(rowid int) *domain.Subscribe {
	all := this.readAll()

	for _, sub := range all {
		if sub.ID == rowid {
			return sub
		}
	}

	return nil
}

func (this JsonRepository) write(newData []*domain.Subscribe) error {
	buff, err := json.Marshal(newData)

	if err != nil {
		return err
	}

	err = os.WriteFile(subscribes_db_file_path, buff, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
