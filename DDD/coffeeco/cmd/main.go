package main

import (
	"context"
	"log"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"

	coffeeco "github.com/ARUMANDESU/Go/DDD/coffeeco/internal"
	"github.com/ARUMANDESU/Go/DDD/coffeeco/internal/payment"
	"github.com/ARUMANDESU/Go/DDD/coffeeco/internal/purchase"
	"github.com/ARUMANDESU/Go/DDD/coffeeco/internal/store"
)

func main() {
	ctx := context.Background()

	// This is the test key from Stripe's documentation. Feel free to use it, no charges will actually be made.
	stripeTestAPIKey := "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

	// This is a test token from Stripe's documentation. Feel free to use it, no charges will actually be made.
	cardToken := "tok_visa"

	// This is the credentials for mongo if you run docker-compose up in this repo.
	mongoConString := "mongodb://root:example@localhost:27017"

	stripeSvc, err := payment.NewStripeService(stripeTestAPIKey)
	if err != nil {
		panic(err)
	}

	purchaseRepo, err := purchase.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err := purchaseRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	storeRepo, err := store.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err := storeRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	storeSvc := store.NewService(storeRepo)

	purchaseSvc := purchase.NewService(stripeSvc, storeSvc, purchaseRepo)

	// This is a test purchase. Feel free to change it.
	pur := &purchase.Purchase{
		CardToken: &cardToken,
		Store: store.Store{
			ID: uuid.New(),
		},
		ProductsToPurchase: []coffeeco.Product{{
			ItemName:  "Coffee",
			BasePrice: *money.New(500, "USD"), // $5.00
			Size:      coffeeco.PRODUCT_SIZE_LARGE,
		}},
		PaymentMeans: payment.MEANS_CARD,
	}
	if err := purchaseSvc.CompletePurchase(ctx, pur, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("purchase was successful")
}
