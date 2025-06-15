package errorf

import (
	"time"
)

// Enum for error codes
type ErrorCode string

const (
	FormaPagamentoInvalida   ErrorCode = "FORMA_PAGAMENTO_INVALIDA"
	DataPagamentoInvalida    ErrorCode = "DATA_PAGAMENTO_INVALIDA"
	DetalhePagamentoInvalido ErrorCode = "DETALHE_PAGAMENTO_INVALIDO"
	ParametroNaoInformado    ErrorCode = "PARAMETRO_NAO_INFORMADO"
	ParametroInvalido        ErrorCode = "PARAMETRO_INVALIDO"
	ErroIdempotencia         ErrorCode = "ERRO_IDEMPOTENCIA"
	NaoInformado             ErrorCode = "NAO_INFORMADO"
)

// ErrorDetail represents a single error
type ErrorDetail struct {
	Code   ErrorCode `json:"code" validate:"required,oneof=FORMA_PAGAMENTO_INVALIDA DATA_PAGAMENTO_INVALIDA DETALHE_PAGAMENTO_INVALIDO PARAMETRO_NAO_INFORMADO PARAMETRO_INVALIDO ERRO_IDEMPOTENCIA NAO_INFORMADO"`
	Title  string    `json:"title" validate:"required,max=255"`
	Detail string    `json:"detail" validate:"required,max=2048"`
}

// Meta holds metadata about the request
type Meta struct {
	RequestDateTime time.Time `json:"requestDateTime" validate:"required"`
}

// ErrorResponse represents the whole response with errors and metadata
type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors" validate:"required,min=1,max=3,dive"` // minItems: 1, maxItems: 3
	Meta   Meta          `json:"meta" validate:"required"`
}

func BuildErrorResponse(code ErrorCode, title string, detail string) ErrorResponse {
	return ErrorResponse{
		Errors: []ErrorDetail{
			{
				Code:   code,
				Title:  title,
				Detail: detail,
			},
		},
		Meta: Meta{
			RequestDateTime: time.Now(),
		},
	}
}
