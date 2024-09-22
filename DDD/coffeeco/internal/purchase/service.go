package purchase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"

	"github.com/ARUMANDESU/Go/DDD/coffeeco/internal/loyalty"
	"github.com/ARUMANDESU/Go/DDD/coffeeco/internal/payment"
	"github.com/ARUMANDESU/Go/DDD/coffeeco/internal/store"
)

type CardChargeService interface {
	ChargeCard(ctx context.Context, amount money.Money, cardToken string) error
}

type StoreService interface {
	GetStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID) (int64, error)
}

type Service struct {
	cardService  CardChargeService
	storeService StoreService
	purchaseRepo Repository
}

func NewService(cardService CardChargeService, storeService StoreService, purchaseRepo Repository) Service {
	return Service{
		cardService:  cardService,
		storeService: storeService,
		purchaseRepo: purchaseRepo,
	}
}

func (s Service) CompletePurchase(ctx context.Context, purchase *Purchase, coffeeBuxCard *loyalty.CoffeeBux) error {
	if err := purchase.validateAndEnrich(); err != nil {
		return err
	}

	if err := s.calculateStoreSpecificDiscount(ctx, purchase.Store.ID, purchase); err != nil {
		return fmt.Errorf("failed to calculate store specific discount: %w", err)
	}

	switch purchase.PaymentMeans {
	case payment.MEANS_CARD:
		if err := s.cardService.ChargeCard(ctx, purchase.total, *purchase.CardToken); err != nil {
			return errors.New("card charge failed, cancelling purchase")
		}
	case payment.MEANS_CASH:
		// do nothing
	case payment.MEANS_COFFEEBUX:
		if err := coffeeBuxCard.Pay(ctx, purchase.ProductsToPurchase); err != nil {
			return fmt.Errorf("failed to charge loyalty card: %w", err)
		}
	default:
		return errors.New("unknown payment method")
	}

	if err := s.purchaseRepo.Store(ctx, *purchase); err != nil {
		return errors.New("could not store purchase")
	}

	// Note: A stamp is added even if the purchase is charged to a loyalty card.
	if coffeeBuxCard != nil {
		coffeeBuxCard.AddStamp()
	}

	return nil
}

func (s Service) calculateStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID, purchase *Purchase) error {
	discount, err := s.storeService.GetStoreSpecificDiscount(ctx, storeID)
	if err != nil && !errors.Is(err, store.ErrNoDiscount) {
		return fmt.Errorf("failed to get discount: %w", err)
	}

	purchasePrice := purchase.total
	if discount > 0 {
		purchase.total = *purchasePrice.Multiply(int64(100 - discount))
	}
	return nil
}
