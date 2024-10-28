package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Amount int64 `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	intent, err := paymentintent.New(params)
	if err != nil {
		http.Error(w, "Payment processing failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(intent.ClientSecret)
}
