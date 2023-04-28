package DHL_API_lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func NewClient(username string, password string, testMode bool, apiAddr string) (*Client, error) {

	client := &Client{
		Username:   username,
		Password:   password,
		TestMode:   testMode,
		APIAddress: apiAddr,
	}

	return client, nil
}

func (c *Client) ValidateAddress(addrString string) (res bool, uniformAddr string, err error) {

	address := ParseAddress(addrString)

	url := "https://api-mock.dhl.com/mydhlapi/address-validate"

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(c.Username, c.Password) // заменить на реальные данные авторизации
	req.Header.Add("type", "delivery")
	req.Header.Add("countryCode", address.CountryCode)
	req.Header.Add("postalCode", address.PostalCode)
	req.Header.Add("cityName", address.CityName)

	response, _ := http.DefaultClient.Do(req)

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response)
	fmt.Println(string(body))
	var validationResponse AddressValidationResponse

	err = json.Unmarshal(body, &validationResponse)
	if err != nil {
		return false, "", err
	}
	for _, addr := range validationResponse.Addresses {
		return true, addr.PostalCode + " " + addr.CityName + " " + addr.CountryCode, nil
	}

	return false, "", errors.New("no such address")
}

func (c *Client) GetStatus(orderID string) (status string, err error) {

	url := c.APIAddress + "/" + orderID + "/tracking"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.Username, c.Password) // заменить на реальные данные авторизации
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("shipmentTrackingNumber", orderID)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response CheckResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if len(response.Shipments) == 0 {
		return "", fmt.Errorf("no shipments found for order with uuid %s", orderID)
	}

	status = response.Shipments[0].Status
	return status, nil
}

func (c *Client) CreateOrder(addrFrom string, addrTo string, size Package, typeSending int) (orderID string, err error) {

	request := CreateOrderRequest{
		PlannedShippingDateAndTime: "2022-10-19T19:19:40 GMT+00:00",
		Pickup: Pickup{
			IsRequested: false,
		},
		ProductCode:      "P",
		LocalProductCode: "P",
		GetRateEstimates: false,
		Accounts: []Account{
			Account{
				TypeCode: "shipper",
				Number:   "123456789",
			},
		},
		ValueAddedServices: []ValueAddedService{
			ValueAddedService{
				ServiceCode: "II",
				Value:       10,
				Currency:    "USD",
			},
		},
		OutputImageProperties: OutputImageProperties{
			PrinterDPI:     300,
			EncodingFormat: "pdf",
			ImageOptions: []ImageOption{
				ImageOption{
					TypeCode:            "invoice",
					TemplateName:        "COMMERCIAL_INVOICE_P_10",
					IsRequested:         true,
					InvoiceType:         "commercial",
					LanguageCode:        "eng",
					LanguageCountryCode: "US",
				},
				ImageOption{
					TypeCode:          "waybillDoc",
					TemplateName:      "ARCH_8x4",
					IsRequested:       true,
					HideAccountNumber: false,
					NumberOfCopies:    1,
				},
				ImageOption{
					TypeCode:      "label",
					TemplateName:  "ECOM26_84_001",
					RenderDHLLogo: true,
					FitLabelsToA4: false,
				},
			},
			SplitTransportAndWaybillDocLabels: true,
			AllDocumentsInOneImage:            false,
			SplitDocumentsByPages:             false,
			SplitInvoiceAndReceipt:            true,
			ReceiptAndLabelsInOneImage:        false,
		},
		CustomerDetails: CustomerDetails{
			ShipperDetails: Details{
				PostalAddress: PostalAddress{
					PostalCode:   "526238",
					CityName:     "Zhaoqing",
					CountryCode:  "CN",
					AddressLine1: "4FENQU, 2HAOKU, WEIPINHUI WULIU YUAN，DAWANG",
					CountryName:  "CHINA, PEOPLES REPUBLIC",
				},
				ContactInformation: ContactInformation{
					Email:    "sender@example.com",
					Phone:    "555-1234",
					FullName: "Sender Name",
				},
				TypeCode: "shipper",
			},
			ReceiverDetails: Details{
				PostalAddress: PostalAddress{
					PostalCode:   "54321",
					CityName:     "Los Angeles",
					CountryCode:  "US",
					AddressLine1: "116 Marine Dr",
					CountryName:  "UNITED STATES OF AMERICA",
				},
				ContactInformation: ContactInformation{
					Email:    "recipient@example.com",
					Phone:    "555-5678",
					FullName: "Recipient Name",
				},
				TypeCode: "recipient",
			},
		},
		Content: Content{
			Packages: []Packages{
				{
					TypeCode: "2BP",
					Weight:   float64(size.Weight),
					Dimensions: Dimensions{
						Length: float64(size.Length),
						Width:  float64(size.Width),
						Height: float64(size.Height),
					},
					CustomerReferences: []Codes{
						{
							Value:    "234546",
							TypeCode: "CU",
						},
					},
					Description:      "description",
					LabelDescription: "bespoke label description",
				},
			},
			IsCustomsDeclarable:   true,
			DeclaredValue:         120,
			DeclaredValueCurrency: "USD",
			ExportDeclaration: ExportDeclaration{
				LineItems: []LineItems{
					{
						Number:      1,
						Description: "bla bla",
						Price:       15,
						Quantity: Quantity{
							Value:             2,
							UnitOfMeasurement: "GM",
						},
						CommodityCodes: []Codes{
							{
								TypeCode: "outbound",
								Value:    "84713000",
							},
							{
								TypeCode: "inbound",
								Value:    "5109101110",
							},
						},
						ExportReasonType:                  "permanent",
						ManufacturerCountry:               "US",
						ExportControlClassificationNumber: "US123456789",
						Weight: Weight{
							NetValue:   0.1,
							GrossValue: float64(size.Weight),
						},
						IsTaxesPaid: true,
						CustomerReferences: []Codes{
							{
								TypeCode: "AFE",
								Value:    "1299210",
							},
						},
						CustomsDocuments: []Codes{
							{
								TypeCode: "COO",
								Value:    "MyDHLAPI - LN#1-CUSDOC-001",
							},
						},
					},
				},
				Invoice: Invoice{
					Number:           "2667168671",
					Date:             "2022-10-22",
					Instructions:     []string{"Handle with care"},
					TotalNetWeight:   0.4,
					TotalGrossWeight: float64(size.Weight),
					CustomerReferences: []Codes{
						{
							TypeCode: "UCN",
							Value:    "UCN-783974937",
						},
						{
							TypeCode: "CN",
							Value:    "CUN-76498376498",
						},
						{
							TypeCode: "RMA",
							Value:    "MyDHLAPI-TESTREF-001",
						},
					},
					TermsOfPayment: "100 days",
					IndicativeCustomsValues: IndicativeCustomsValues{
						ImportCustomsDutyValue: 150.57,
						ImportTaxesValue:       49.43,
					},
				},
				AdditionalCharges: []AdditionalCharges{
					{
						Value:    10,
						Caption:  "fee",
						TypeCode: "freight",
					},
				},
				DestinationPortName: "New Yourk Port",
				PlaceOfIncoterm:     "ShenZhen Port",
				PayerVATNumber:      "12345ED",
				RecipientReference:  "01291344",
				Exporter: Exporter{
					Id:   "121233",
					Code: "S",
				},
				PackageMarks:     "Fragile glass bottle",
				ExportReference:  "export referance",
				ExportReason:     "er",
				ExportReasonType: "permanent",
				Licenses: []Codes{
					{
						TypeCode: "export",
						Value:    "123127233",
					},
				},
				ShipmentType: "personal",
				CustomsDocuments: []Codes{
					{
						TypeCode: "INV",
						Value:    "MyDHLAPI - CUSDOC-001",
					},
				},
			},
			Description:       "Shipment",
			USFilingTypeValue: "12345",
			Incoterm:          "DAP",
			UnitOfMeasurement: "metric",
		},
		ShipmentNotification: []ShipmentNotification{
			{
				TypeCode:            "email",
				ReceiverId:          "shipmentnotification@mydhlapisample.com",
				LanguageCode:        "eng",
				LanguageCountryCode: "UK",
				BespokeMessage:      "message to be included in the notification",
			},
		},
		GetAdditionalInformation: []EstimatedDeliveryDate{
			{
				TypeCode:    "pickupDetails",
				IsRequested: true,
			},
		},
	}

	data, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.APIAddress, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result OrderCreationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if result.ShipmentTrackingNumber == "" {

		return "", errors.New("can't create shipement")
	}

	return result.ShipmentTrackingNumber, nil
}

func ParseAddress(addrString string) Address {
	var address Address
	// Разделение строки на составляющие
	parts := strings.FieldsFunc(addrString, func(r rune) bool {
		return r == ',' || r == ' '
	})

	// Определение индекса, названия города и кода страны
	for _, part := range parts {
		if len(part) == 2 && part == strings.ToUpper(part) {
			address.CountryCode = part
		} else if len(part) == 5 && strings.HasSuffix(part, "00") {
			address.PostalCode = part
		} else {
			address.CityName += part + " "
		}
	}
	address.CityName = strings.TrimSpace(address.CityName)

	return address
}

func (c *Client) Calculate(addrFrom string, addrTo string, size Package) ([]PriceSending, error) {
	// Определяем параметры запроса
	req, err := http.NewRequest("GET", c.APIAddress, nil)
	if err != nil {
		return nil, err
	}
	origin := ParseAddress(addrFrom)
	destination := ParseAddress(addrTo)

	req.SetBasicAuth(c.Username, c.Password) // заменить на реальные данные авторизации
	q := req.URL.Query()
	q.Add("accountNumber", "123456789")
	q.Add("originCountryCode", origin.CountryCode)
	q.Add("originPostalCode", origin.PostalCode)
	q.Add("originCityName", origin.CityName)
	q.Add("destinationCountryCode", destination.CountryCode)
	q.Add("destinationPostalCode", destination.PostalCode)
	q.Add("destinationCityName", destination.CityName)
	q.Add("weight", fmt.Sprintf("%v", size.Weight))
	q.Add("length", fmt.Sprintf("%v", size.Length))
	q.Add("width", fmt.Sprintf("%v", size.Width))
	q.Add("height", fmt.Sprintf("%v", size.Height))
	q.Add("plannedShippingDate", "2020-02-26")
	q.Add("isCustomsDeclarable", "false")
	q.Add("unitOfMeasurement", "metric")
	q.Add("nextBusinessDay", "false")
	q.Add("strictValidation", "false")
	q.Add("getAllValueAddedServices", "false")
	q.Add("requestEstimatedDeliveryDate", "true")
	q.Add("estimatedDeliveryDateType", "QDDF")
	req.URL.RawQuery = q.Encode()

	// Отправляем запрос и обрабатываем ответ
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}
	// Чтение ответа
	var result CalculateRequestBody

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	prices := []PriceSending{}
	for _, product := range result.Products {
		for _, price := range product.TotalPrices {
			prices = append(prices, PriceSending{
				ProductName:        product.ProductName,
				TotalPriceCurrency: price.CurrencyType,
				TotalPriceValue:    price.Price,
			})
		}
	}

	return prices, nil
}
