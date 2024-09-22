package store

import (
	"github.com/google/uuid"

	coffeeco "github.com/ARUMANDESU/Go/DDD/coffeeco/internal"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeeco.Product
}
