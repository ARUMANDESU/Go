package loyalty

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	coffeeco "github.com/ARUMANDESU/Go/DDD/coffeeco/internal"
	"github.com/ARUMANDESU/Go/DDD/coffeeco/internal/store"
)

const DefaultDrinksUntilFree = 10

type CoffeeBux struct {
	ID                  uuid.UUID
	Store               store.Store
	coffeeLover         coffeeco.CoffeeLover
	FreeDrinksAvailable int
	// The number of drinks a customer needs to purchase before they get a free drink
	DrinksUntilFree int
}

func (c *CoffeeBux) AddStamp() {
	if c.DrinksUntilFree == 1 {
		c.DrinksUntilFree = DefaultDrinksUntilFree
		c.FreeDrinksAvailable += 1
	} else {
		c.DrinksUntilFree--
	}
}

func (c *CoffeeBux) Pay(ctx context.Context, purchases []coffeeco.Product) error {
	lp := len(purchases)
	if lp == 0 {
		return errors.New("nothing to buy")
	}

	if c.FreeDrinksAvailable < lp {
		return fmt.Errorf("not enough coffeeBux to cover entire purchase. Have %d, need %d", len(purchases), c.FreeDrinksAvailable)
	}

	c.FreeDrinksAvailable = c.FreeDrinksAvailable - lp
	return nil
}
