package entity

type OrderDetail struct {
	ID         int64
	OrderID    int64
	ProductID  string
	Quantity   float64
	Price      float64
	TotalPrice float64
}
