package entity


type MerchantAlias struct {
	DBAName string `xml:"DBAName"`
	LegalName string `xml:"LegalName"`
	CorporateName string `xml:"CorporateName"`
	BASEIIName string `xml:"BASEIIName"`
}