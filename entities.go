package DHL_API_lib

type PriceSending struct {
	ProductName            string  `json:"productName"`
	TotalPriceCurrencyType string  `json:"totalPriceCurrencyType"`
	TotalPriceCurrency     string  `json:"totalPriceCurrency"`
	TotalPriceValue        float64 `json:"totalPriceValue"`
}

type Client struct {
	Username   string
	Password   string
	TestMode   bool
	APIAddress string
}

type Address struct {
	CountryCode string `json:"countryCode"`
	PostalCode  string `json:"postalCode"`
	CityName    string `json:"cityName"`
	CountyName  string `json:"countyName"`
}

type CalculateRequestBody struct {
	Products []Product `json:"products"`
}

type Product struct {
	ProductName string       `json:"productName"`
	TotalPrices []TotalPrice `json:"totalPrice"`
}

type TotalPrice struct {
	CurrencyType  string  `json:"currencyType"`
	PriceCurrency string  `json:"priceCurrency"`
	Price         float64 `json:"price"`
}

type AddressValidationResponse struct {
	Warnings  []string   `json:"warnings"`
	Addresses []*Address `json:"address"`
}

type CreateOrderRequest struct {
	PlannedShippingDateAndTime string                  `json:"plannedShippingDateAndTime"`
	Pickup                     Pickup                  `json:"pickup"`
	ProductCode                string                  `json:"productCode"`
	LocalProductCode           string                  `json:"localProductCode"`
	GetRateEstimates           bool                    `json:"getRateEstimates"`
	Accounts                   []Account               `json:"accounts"`
	ValueAddedServices         []ValueAddedService     `json:"valueAddedServices"`
	OutputImageProperties      OutputImageProperties   `json:"outputImageProperties"`
	CustomerDetails            CustomerDetails         `json:"customerDetails"`
	Content                    Content                 `json:"content"`
	ShipmentNotification       []ShipmentNotification  `json:"shipmentNotification"`
	GetTransliteratedResponse  bool                    `json:"getTransliteratedResponse"`
	EstimatedDeliveryDate      EstimatedDeliveryDate   `json:"estimatedDeliveryDate"`
	GetAdditionalInformation   []EstimatedDeliveryDate `json:"getAdditionalInformation"`
}

type EstimatedDeliveryDate struct {
	IsRequested bool   `json:"isRequested"`
	TypeCode    string `json:"typeCode"`
}

type ShipmentNotification struct {
	TypeCode            string `json:"typeCode"`
	ReceiverId          string `json:"receiverId"`
	LanguageCode        string `json:"languageCode"`
	LanguageCountryCode string `json:"languageCountryCode"`
	BespokeMessage      string `json:"bespokeMessage"`
}

type Pickup struct {
	IsRequested bool `json:"isRequested"`
}

type Account struct {
	TypeCode string `json:"typeCode"`
	Number   string `json:"number"`
}

type ValueAddedService struct {
	ServiceCode string `json:"serviceCode"`
	Value       int    `json:"value"`
	Currency    string `json:"currency"`
}

type OutputImageProperties struct {
	PrinterDPI                        int           `json:"printerDPI"`
	EncodingFormat                    string        `json:"encodingFormat"`
	ImageOptions                      []ImageOption `json:"imageOptions"`
	SplitTransportAndWaybillDocLabels bool          `json:"splitTransportAndWaybillDocLabels"`
	AllDocumentsInOneImage            bool          `json:"allDocumentsInOneImage"`
	SplitDocumentsByPages             bool          `json:"splitDocumentsByPages"`
	SplitInvoiceAndReceipt            bool          `json:"splitInvoiceAndReceipt"`
	ReceiptAndLabelsInOneImage        bool          `json:"receiptAndLabelsInOneImage"`
}

type ImageOption struct {
	TypeCode            string `json:"typeCode"`
	TemplateName        string `json:"templateName"`
	IsRequested         bool   `json:"isRequested"`
	InvoiceType         string `json:"invoiceType,omitempty"`
	LanguageCode        string `json:"languageCode,omitempty"`
	LanguageCountryCode string `json:"languageCountryCode,omitempty"`
	HideAccountNumber   bool   `json:"hideAccountNumber,omitempty"`
	NumberOfCopies      int    `json:"numberOfCopies,omitempty"`
	RenderDHLLogo       bool   `json:"renderDHLLogo,omitempty"`
	FitLabelsToA4       bool   `json:"fitLabelsToA4,omitempty"`
}

type CustomerDetails struct {
	ShipperDetails  Details `json:"shipperDetails"`
	ReceiverDetails Details `json:"receiverDetails"`
}

type Details struct {
	PostalAddress      PostalAddress      `json:"postalAddress"`
	ContactInformation ContactInformation `json:"contactInformation"`
	TypeCode           string             `json:"typeCode"`
}

type PostalAddress struct {
	PostalCode   string `json:"postalCode,omitempty"`
	CityName     string `json:"cityName"`
	CountryCode  string `json:"countryCode"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	AddressLine3 string `json:"addressLine3,omitempty"`
	CountyName   string `json:"countyName,omitempty"`
	CountryName  string `json:"countryName"`
}

type ContactInformation struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	FullName    string `json:"fullName"`
}

type PhoneNumber struct {
	Number string `json:"number"`
}

type Package struct {
	Weight int `json:"weight"`
	Length int `json:"length"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type OrderCreationResponse struct {
	ShipmentTrackingNumber string                `json:"shipmentTrackingNumber"`
	TrackingUrl            string                `json:"trackingUrl"`
	Packages               []DHLPackage          `json:"packages"`
	Documents              []Document            `json:"documents"`
	ShipmentDetails        []ShipmentDetail      `json:"shipmentDetails"`
	EstimatedDeliveryDate  EstimatedDeliveryDate `json:"estimatedDeliveryDate"`
}

type DHLPackage struct {
	ReferenceNumber int    `json:"referenceNumber"`
	TrackingNumber  string `json:"trackingNumber"`
	TrackingUrl     string `json:"trackingUrl"`
}

type Document struct {
	ImageFormat string `json:"imageFormat"`
	Content     string `json:"content"`
	TypeCode    string `json:"typeCode"`
}

type ShipmentDetail struct {
	PickupDetails PickupDetails `json:"pickupDetails"`
}

type PickupDetails struct {
	LocalCutoffDateAndTime string `json:"localCutoffDateAndTime"`
	GmtCutoffTime          string `json:"gmtCutoffTime"`
	CutoffTimeOffset       string `json:"cutoffTimeOffset"`
	PickupEarliest         string `json:"pickupEarliest"`
	PickupLatest           string `json:"pickupLatest"`
	TotalTransitDays       string `json:"totalTransitDays"`
	PickupAdditionalDays   string `json:"pickupAdditionalDays"`
	DeliveryAdditionalDays string `json:"deliveryAdditionalDays"`
	PickupDayOfWeek        string `json:"pickupDayOfWeek"`
	DeliveryDayOfWeek      string `json:"deliveryDayOfWeek"`
}

type Content struct {
	Packages              []Packages        `json:"packages"`
	IsCustomsDeclarable   bool              `json:"isCustomsDeclarable"`
	DeclaredValue         float64           `json:"declaredValue"`
	DeclaredValueCurrency string            `json:"declaredValueCurrency"`
	ExportDeclaration     ExportDeclaration `json:"exportDeclaration"`
	Description           string            `json:"description"`
	USFilingTypeValue     string            `json:"USFilingTypeValue"`
	Incoterm              string            `json:"incoterm"`
	UnitOfMeasurement     string            `json:"unitOfMeasurement"`
}

type Packages struct {
	TypeCode           string     `json:"typeCode"`
	Weight             float64    `json:"weight"`
	Dimensions         Dimensions `json:"dimensions"`
	CustomerReferences []Codes    `json:"customerReferences"`
	Description        string     `json:"description"`
	LabelDescription   string     `json:"labelDescription"`
}

type Dimensions struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type ExportDeclaration struct {
	LineItems []LineItems `json:"lineItems"`
	Invoice   Invoice     `json:"invoice"`
	Remarks   []struct {
		Value string `json:"value"`
	} `json:"remarks"`
	AdditionalCharges   []AdditionalCharges `json:"additionalCharges"`
	DestinationPortName string              `json:"destinationPortName"`
	PlaceOfIncoterm     string              `json:"placeOfIncoterm"`
	PayerVATNumber      string              `json:"payerVATNumber"`
	RecipientReference  string              `json:"recipientReference"`
	Exporter            Exporter            `json:"exporter"`
	PackageMarks        string              `json:"packageMarks"`
	DeclarationNotes    []struct {
		Value string `json:"value"`
	} `json:"declarationNotes"`
	ExportReference  string  `json:"exportReference"`
	ExportReason     string  `json:"exportReason"`
	ExportReasonType string  `json:"exportReasonType"`
	Licenses         []Codes `json:"licenses"`
	ShipmentType     string  `json:"shipmentType"`
	CustomsDocuments []Codes `json:"customsDocuments"`
}

type Exporter struct {
	Id   string `json:"id"`
	Code string `json:"code"`
}

type AdditionalCharges struct {
	Value    int    `json:"value"`
	Caption  string `json:"caption"`
	TypeCode string `json:"typeCode"`
}

type Invoice struct {
	Number                  string                  `json:"number"`
	Date                    string                  `json:"date"`
	Instructions            []string                `json:"instructions"`
	TotalNetWeight          float64                 `json:"totalNetWeight"`
	TotalGrossWeight        float64                 `json:"totalGrossWeight"`
	CustomerReferences      []Codes                 `json:"customerReferences"`
	TermsOfPayment          string                  `json:"termsOfPayment"`
	IndicativeCustomsValues IndicativeCustomsValues `json:"indicativeCustomsValues"`
}

type IndicativeCustomsValues struct {
	ImportCustomsDutyValue float64 `json:"importCustomsDutyValue"`
	ImportTaxesValue       float64 `json:"importTaxesValue"`
}

type LineItems struct {
	Number                            int64    `json:"number"`
	Description                       string   `json:"description"`
	Price                             float64  `json:"price"`
	Quantity                          Quantity `json:"quantity"`
	CommodityCodes                    []Codes  `json:"commodityCodes"`
	ExportReasonType                  string   `json:"exportReasonType"`
	ManufacturerCountry               string   `json:"manufacturerCountry"`
	ExportControlClassificationNumber string   `json:"exportControlClassificationNumber"`
	Weight                            Weight   `json:"weight"`
	IsTaxesPaid                       bool     `json:"isTaxesPaid"`
	AdditionalInformation             []string `json:"additionalInformation"`
	CustomerReferences                []Codes  `json:"customerReferences"`
	CustomsDocuments                  []Codes  `json:"customsDocuments"`
}

type Weight struct {
	NetValue   float64 `json:"netValue"`
	GrossValue float64 `json:"grossValue"`
}

type Codes struct {
	TypeCode string `json:"typeCode"`
	Value    string `json:"value"`
}

type Quantity struct {
	Value             int    `json:"value"`
	UnitOfMeasurement string `json:"unitOfMeasurement"`
}

type CheckResponse struct {
	Shipments []Shipment
}

type Shipment struct {
	ShipmentTrackingNumber string         `json:"shipmentTrackingNumber"`
	Status                 string         `json:"status"`
	ShipmentTimestamp      string         `json:"shipmentTimestamp"`
	ProductCode            string         `json:"productCode"`
	Description            string         `json:"description"`
	ShipperDetails         ShipperDetails `json:"shipperDetails"`
	Events                 []Events       `json:"events"`
	EstimatedDeliveryDate  string         `json:"estimatedDeliveryDate"`
}

type Events struct {
	Date        string `json:"date"`
	Time        string `json:"time"`
	TypeCode    string `json:"typeCode"`
	Description string `json:"description"`
}
type ShipperDetails struct {
	Name            string          `json:"name"`
	PostalAddress   PostalAddress   `json:"postalAddress"`
	ReceiverDetails ReceiverDetails `json:"receiverDetails"`
}

type ReceiverDetails struct {
	Name          string        `json:"name"`
	PostalAddress PostalAddress `json:"postalAddress"`
}
