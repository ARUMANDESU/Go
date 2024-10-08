package purchase

import (
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"

	coffeeco "github.com/ARUMANDESU/Go/DDD/coffeeco/internal"
	"github.com/ARUMANDESU/Go/DDD/coffeeco/internal/payment"
	"github.com/ARUMANDESU/Go/DDD/coffeeco/internal/store"
)

type Purchase struct {
	id                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeeco.Product
	total              money.Money
	PaymentMeans       payment.Means
	timeOfPurchase     time.Time
	CardToken          *string
}

func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductsToPurchase) == 0 {
		return errors.New("purchease must contain at least one product")
	}

	p.total = *money.New(0, "USD")

	for _, product := range p.ProductsToPurchase {
		newTotal, _ := p.total.Add(&product.BasePrice)
		p.total = *newTotal
	}
	if p.total.IsZero() {
		return errors.New("likely mistake; purchase should never be 0. Please validate")
	}

	p.id = uuid.New()
	p.timeOfPurchase = time.Now()

	return nil
}
