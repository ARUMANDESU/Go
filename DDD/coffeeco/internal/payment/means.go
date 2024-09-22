package payment

type Means string

const (
	MEANS_CASH      Means = "cash"
	MEANS_CARD      Means = "card"
	MEANS_COFFEEBUX Means = "coffeebux"
)

type CardDetails struct {
	cardToken string
}
