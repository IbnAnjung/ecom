package order

import "fmt"

func (uc *orderUsecase) CancelExpiredOrder() {
	fmt.Println("from usecase")
}
