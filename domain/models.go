package domain

type Subscribe struct {
	ID     int    `json:"id"`
	ShopID string `json:"shopId"`
	Url    string `json:"productUrl"`
	Name   string `json:"productName"`
}

type Monitoring struct {
	SubscribeID int     `json:"subscribeId"`
	Timestamp   string  `json:"timestamp"`
	Price       float64 `json:"price"`
}

type CreateNewRequest struct {
	ProductId string `json:"productId"`
	Link      string `json:"link"`
	Name      string `json:"name"`
}
