package web

type PaymentMidtrans struct {
	Token       string `json:"token"`
	RedirectUrl string `json:"redirect_url"`
}
