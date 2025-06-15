package consent

import (
	"context"
	"openfinance/internal/core/domain/consent"
)

type AdapterConsentService struct{}

func SaveConsent(ctx context.Context, request consent.CreateConsentRequest, xIdempotencyKey string) (consent.CreateConsentResponse, error) {
	return consent.CreateConsentResponse{}, nil
}
