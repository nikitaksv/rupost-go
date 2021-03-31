package otpravka

import (
	"context"
	"net/http"
)

type OrderService service

type OrderSearchResponse struct {
	Orders []Order `json:"orders"`
}

func (s OrderService) Search(ctx context.Context, query string) (*OrderSearchResponse, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/1.0/backlog/search", ContentTypeJSON, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	c := new(OrderSearchResponse)
	c.Orders = []Order{}
	resp, err := s.client.Do(ctx, req, &c.Orders)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

type Order struct {
	AddressChanged                  bool               `json:"address-changed"`
	AddressTypeTo                   string             `json:"address-type-to"`
	AreaTo                          string             `json:"area-to"`
	AviaRate                        int                `json:"avia-rate"`
	AviaRateWithVat                 int                `json:"avia-rate-with-vat"`
	AviaRateWoVat                   int                `json:"avia-rate-wo-vat"`
	Barcode                         string             `json:"barcode"`
	BrandName                       string             `json:"brand-name"`
	BuildingTo                      string             `json:"building-to"`
	Comment                         string             `json:"comment"`
	CompletenessChecking            bool               `json:"completeness-checking"`
	CompletenessCheckingRateWithVat int                `json:"completeness-checking-rate-with-vat"`
	CompletenessCheckingRateWoVat   int                `json:"completeness-checking-rate-wo-vat"`
	CorpusTo                        string             `json:"corpus-to"`
	CustomsDeclaration              CustomsDeclaration `json:"customs-declaration"`
	DeliveryTime                    DeliveryTime       `json:"delivery-time"`
	DeliveryWithCod                 bool               `json:"delivery-with-cod"`
	Dimension                       Dimension          `json:"dimension"`
	EnvelopeType                    string             `json:"envelope-type"`
	FragileRateWithVat              int                `json:"fragile-rate-with-vat"`
	FragileRateWoVat                int                `json:"fragile-rate-wo-vat"`
	GivenName                       string             `json:"given-name"`
	Goods                           Goods              `json:"goods"`
	GroundRate                      int                `json:"ground-rate"`
	GroundRateWithVat               int                `json:"ground-rate-with-vat"`
	GroundRateWoVat                 int                `json:"ground-rate-wo-vat"`
	HotelTo                         string             `json:"hotel-to"`
	HouseTo                         string             `json:"house-to"`
	ID                              int                `json:"id"`
	IndexTo                         int                `json:"index-to"`
	InsrRate                        int                `json:"insr-rate"`
	InsrRateWithVat                 int                `json:"insr-rate-with-vat"`
	InsrRateWoVat                   int                `json:"insr-rate-wo-vat"`
	InsrValue                       int                `json:"insr-value"`
	InventoryRateWithVat            int                `json:"inventory-rate-with-vat"`
	InventoryRateWoVat              int                `json:"inventory-rate-wo-vat"`
	IsDeleted                       bool               `json:"is-deleted"`
	LetterTo                        string             `json:"letter-to"`
	LocationTo                      string             `json:"location-to"`
	MailCategory                    string             `json:"mail-category"`
	MailDirect                      int                `json:"mail-direct"`
	MailRank                        string             `json:"mail-rank"`
	MailType                        string             `json:"mail-type"`
	ManualAddressInput              bool               `json:"manual-address-input"`
	Mass                            int                `json:"mass"`
	MassRate                        int                `json:"mass-rate"`
	MassRateWithVat                 int                `json:"mass-rate-with-vat"`
	MassRateWoVat                   int                `json:"mass-rate-wo-vat"`
	MiddleName                      string             `json:"middle-name"`
	NoticePaymentMethod             string             `json:"notice-payment-method"`
	NoticeRateWithVat               int                `json:"notice-rate-with-vat"`
	NoticeRateWoVat                 int                `json:"notice-rate-wo-vat"`
	NumAddressTypeTo                string             `json:"num-address-type-to"`
	OfficeTo                        string             `json:"office-to"`
	OrderNum                        string             `json:"order-num"`
	OversizeRateWithVat             int                `json:"oversize-rate-with-vat"`
	OversizeRateWoVat               int                `json:"oversize-rate-wo-vat"`
	Payment                         int                `json:"payment"`
	PaymentMethod                   string             `json:"payment-method"`
	PlaceTo                         string             `json:"place-to"`
	Postmarks                       []string           `json:"postmarks"`
	PostofficeCode                  string             `json:"postoffice-code"`
	RawAddress                      string             `json:"raw-address"`
	RecipientName                   string             `json:"recipient-name"`
	RegionTo                        string             `json:"region-to"`
	RoomTo                          string             `json:"room-to"`
	SlashTo                         string             `json:"slash-to"`
	SmsNoticeRecipient              int                `json:"sms-notice-recipient"`
	SmsNoticeRecipientRateWithVat   int                `json:"sms-notice-recipient-rate-with-vat"`
	SmsNoticeRecipientRateWoVat     int                `json:"sms-notice-recipient-rate-wo-vat"`
	StrIndexTo                      string             `json:"str-index-to"`
	StreetTo                        string             `json:"street-to"`
	Surname                         string             `json:"surname"`
	TelAddress                      int                `json:"tel-address"`
	TotalRateWoVat                  int                `json:"total-rate-wo-vat"`
	TotalVat                        int                `json:"total-vat"`
	TransportMode                   string             `json:"transport-mode"`
	TransportType                   string             `json:"transport-type"`
	Version                         int                `json:"version"`
	VladenieTo                      string             `json:"vladenie-to"`
}

type CustomsEntries struct {
	Amount      int    `json:"amount"`
	CountryCode int    `json:"country-code"`
	Description string `json:"description"`
	TnvedCode   string `json:"tnved-code"`
	Value       int    `json:"value"`
	Weight      int    `json:"weight"`
}

type CustomsDeclaration struct {
	Currency        string           `json:"currency"`
	CustomsEntries  []CustomsEntries `json:"customs-entries"`
	EntriesType     string           `json:"entries-type"`
	WithCertificate bool             `json:"with-certificate"`
	WithInvoice     bool             `json:"with-invoice"`
	WithLicense     bool             `json:"with-license"`
}

type DeliveryTime struct {
	MaxDays int `json:"max-days"`
	MinDays int `json:"min-days"`
}

type Dimension struct {
	Height int `json:"height"`
	Length int `json:"length"`
	Width  int `json:"width"`
}

type Items struct {
	Description string `json:"description"`
	InsrValue   int    `json:"insr-value"`
	Quantity    int    `json:"quantity"`
	Value       int    `json:"value"`
	VatRate     int    `json:"vat-rate"`
}

type Goods struct {
	Items []Items `json:"items"`
}
