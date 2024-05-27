package entity

type IOrder interface {
	CancelExpiredOrder()
}
