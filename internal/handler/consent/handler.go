package consent

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"openfinance/configuration/logger"
	"openfinance/internal/core/domain/consent"
	"openfinance/internal/core/domain/errorf"
	"openfinance/internal/core/ports"
)

type Handler struct {
	consentService ports.ConsentService
}

func New(consentService ports.ConsentService) *Handler {
	return &Handler{
		consentService: consentService,
	}
}

func (h *Handler) PostConsent(w http.ResponseWriter, r *http.Request) {
	//authorization := r.Header.Get("Authorization")
	//xFapiAuthDate := r.Header.Get("x-fapi-auth-date")
	//xFapiCustomerIpAddress := r.Header.Get("x-fapi-customer-ip-address")
	//xFapiInteractionId := r.Header.Get("x-fapi-interaction-id")
	//xCustomerUserAgent := r.Header.Get("x-customer-user-agent")
	xIdempotencyKey := r.Header.Get("x-idempotency-key")

	var consentRequest consent.CreateConsentRequest

	if err := json.NewDecoder(r.Body).Decode(&consentRequest); err != nil {
		logger.L().Error("Error parsing request body", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorf.BuildErrorResponse(errorf.NaoInformado, "Error Parsing Json", "Error Parsing Json")
		json.NewEncoder(w).Encode(er)
	}

	res, err := h.consentService.SaveConsent(context.Background(), consentRequest, xIdempotencyKey)

	if err != nil {
		logger.L().Error("Save consent service error:", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		er := errorf.BuildErrorResponse(errorf.NaoInformado, "During service layer call", "During service layer call")
		json.NewEncoder(w).Encode(er)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}
