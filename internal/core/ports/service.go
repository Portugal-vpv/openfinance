package ports

import (
	"context"
	"openfinance/internal/core/domain/consent"
)

type ConsentService interface {
	SaveConsent(ctx context.Context, request consent.CreateConsentRequest, xIdempotencyKey string) (consent.CreateConsentResponse, error)
}
