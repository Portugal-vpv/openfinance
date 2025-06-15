package consent

type EnumPaymentType string

const (
	PaymentTypePIX EnumPaymentType = "PIX"
)

type EnumLocalInstrument string

const (
	LocalInstrumentMANU EnumLocalInstrument = "MANU"
	LocalInstrumentDICT EnumLocalInstrument = "DICT"
	LocalInstrumentQRDN EnumLocalInstrument = "QRDN"
	LocalInstrumentQRES EnumLocalInstrument = "QRES"
	LocalInstrumentINIC EnumLocalInstrument = "INIC"
)

type EnumAccountPaymentType string

const (
	AccountTypeCACC EnumAccountPaymentType = "CACC"
	AccountTypeSVGS EnumAccountPaymentType = "SVGS"
	AccountTypeTRAN EnumAccountPaymentType = "TRAN"
)

// EnumPaymentPersonType defines the type of person: individual or legal entity.
type EnumPaymentPersonType string

const (
	PersonTypeNatural EnumPaymentPersonType = "NATURAL"
	PersonTypeLegal   EnumPaymentPersonType = "LEGAL"
)

type CreateConsentRequest struct {
	LoggedUser     *LoggedUser     `json:"loggedUser"`
	BusinessEntity *BusinessEntity `json:"business_entity"`
	Creditor       *Creditor       `json:"creditor"`
	Payment        *Payment        `json:"payment"`
	DebtorAccount  *DebtorAccount  `json:"debtor_account"`
}

type LoggedUser struct {
	Document *Document
}

// Document represents the identification document of a user.
type Document struct {
	Identification *string `json:"identification" validate:"required,len=11,numeric"`
	Rel            *string `json:"rel" validate:"required,len=3,uppercase"`
}

// BusinessEntity represents a legal entity user.
type BusinessEntity struct {
	Document *Document `json:"document" validate:"required"` // Required if user is a legal entity (CNPJ)
}

// Identification holds creditor's identification info.
type Identification struct {
	PersonType *EnumPaymentPersonType `json:"personType" validate:"required,oneof=NATURAL LEGAL"`
	CpfCnpj    *string                `json:"cpfCnpj" validate:"required,min=11,max=14,numeric"`
	Name       *string                `json:"name" validate:"required,max=120,personname"`
}

// Creditor defines the person who receives the payment
type Creditor struct {
	Identification *Identification `json:"identification" validate:"required"`
}

type CreditorAccount struct {
	Ispb        *string                 `json:"ispb" validate:"required,len=8,numeric"`
	Issuer      *string                 `json:"issuer" validate:"omitempty,min=1,max=4,numeric"`
	Number      *string                 `json:"number" validate:"required,min=1,max=20,numeric"`
	AccountType *EnumAccountPaymentType `json:"accountType" validate:"required,oneof=CACC SVGS TRAN"`
}

type Details struct {
	LocalInstrument *EnumLocalInstrument `json:"localInstrument" validate:"required,oneof=MANU DICT QRDN QRES INIC"`
	QRCode          *string              `json:"qrCode,omitempty" validate:"omitempty,max=512"`
	Proxy           *string              `json:"proxy,omitempty" validate:"omitempty,max=77"`
	CreditorAccount *CreditorAccount     `json:"creditorAccount" validate:"required"`
}
type ScheduleSingle struct {
	Date *string `json:"date" validate:"required,len=10,datetime=2006-01-02"`
}

type ScheduleDaily struct {
	StartDate *string `json:"startDate" validate:"required,len=10,datetime=2006-01-02"`
	Quantity  *int    `json:"quantity" validate:"required,min=2,max=60"`
}

type ScheduleWeekly struct {
	DayOfWeek *string `json:"dayOfWeek" validate:"required"` // e.g., "QUINTA_FEIRA"
	StartDate *string `json:"startDate" validate:"required,len=10,datetime=2006-01-02"`
	Quantity  *int    `json:"quantity" validate:"required,min=2,max=60"`
}

type ScheduleMonthly struct {
	DayOfMonth *int    `json:"dayOfMonth" validate:"required,min=1,max=31"`
	StartDate  *string `json:"startDate" validate:"required,len=10,datetime=2006-01-02"`
	Quantity   *int    `json:"quantity" validate:"required,min=2,max=24"`
}

type ScheduleCustom struct {
	Dates          []string `json:"dates" validate:"required,dive,len=10,datetime=2006-01-02,unique"`
	AdditionalInfo *string  `json:"additionalInformation,omitempty" validate:"max=255"`
}

type Schedule struct {
	Single  *ScheduleSingle  `json:"single,omitempty"`
	Daily   *ScheduleDaily   `json:"daily,omitempty"`
	Weekly  *ScheduleWeekly  `json:"weekly,omitempty"`
	Monthly *ScheduleMonthly `json:"monthly,omitempty"`
	Custom  *ScheduleCustom  `json:"custom,omitempty"`
}

type Payment struct {
	Type         *EnumPaymentType `json:"type" validate:"required,oneof=PIX"`
	Schedule     *Schedule        `json:"schedule,omitempty"` // Mutually exclusive with Date
	Date         *string          `json:"date,omitempty" validate:"omitempty,len=10,datetime=2006-01-02"`
	Currency     *string          `json:"currency" validate:"required,len=3,uppercase"`
	Amount       *string          `json:"amount" validate:"required,min=4,max=19,regexp=^\\d{1,16}\\.\\d{2}$"`
	IBGETownCode *string          `json:"ibgeTownCode,omitempty" validate:"omitempty,len=7,numeric"`
	Details      *Details         `json:"details" validate:"required"`
}

type DebtorAccount struct {
	Ispb        *string                 `json:"ispb" validate:"required,len=8,numeric"`
	Issuer      *string                 `json:"issuer,omitempty" validate:"omitempty,min=1,max=4,numeric"`
	Number      *string                 `json:"number" validate:"required,min=1,max=20,numeric"`
	AccountType *EnumAccountPaymentType `json:"accountType" validate:"required,oneof=CACC SVGS TRAN"`
}
