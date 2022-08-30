package entity

type LocationAddress struct {
	AddressToLine1 string `xml:"AddressToLine1"`
	AddressToLine2 string `xml:"AddressToLine2"`
	Street string `xml:"Street"`
	City string `xml:"City"`
	StateProvinceCode string `xml:"StateProvinceCode"`
	PostalCode string `xml:"PostalCode"`
}