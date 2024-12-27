package domain

import "time"

type StatusPaymentMethod string
type StatusPayment string

const (
	MethodMitrans    StatusPaymentMethod = "mitrans"
	PaymentPending   StatusPayment       = "pending"
	PaymentCompleted StatusPayment       = "completed"
	PaymentFailed    StatusPayment       = "failed"
)

type Payment struct {
	Id            int
	PesananId     int
	PaymentMethod StatusPaymentMethod
	PaymentStatus StatusPayment
	PaymentAmount int
	PaymentDate   time.Time
}
