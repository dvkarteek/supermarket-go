package data

import "github.com/google/uuid"

type Item struct {
	Name string `json:"name`
}

type Order struct {
	ID     string   `json:"id"`
	UserID string   `json:"user"`
	Items  []string `json:"items"`
	Status bool     `json:"status"`
}

var orders []*Order = []*Order{}

func StartOrder(u string) *Order {
	o := &Order{
		ID:     uuid.New().String(),
		UserID: u,
		Items:  []string{},
	}
	orders = append(orders, o)
	return o
}

func AddItem(o string, i string) *Order {

	if o == "" || i == "" {
		return nil
	}

	var order *Order
	for _, temp := range orders {
		if temp.ID == o {
			temp.Items = append(temp.Items, i)
			order = temp
			break
		}
	}

	return order

}

func CreateOrder(o string) (int, *Order) {

	if o == "" {
		return -1, nil
	}

	var order *Order = nil
	for _, temp := range orders {
		if temp.ID == o {
			order = temp
			break
		}
	}

	if order == nil {
		return -1, nil
	}
	if order.Status == true {
		return -2, nil
	}

	order.Status = true
	return 0, order

}

func GetOrders(u string) []*Order {

	if u == "" {
		return nil
	}

	var userOrders []*Order
	for _, temp := range orders {
		if temp.UserID == u {
			userOrders = append(userOrders, temp)
		}
	}
	return userOrders
}
