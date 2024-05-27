package entity

import (
	"context"
)

type IInventoryRepository interface {
	GetProductInventories(ctx context.Context, ids []string) (inv []ProductInventory, err error)
}
