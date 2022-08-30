package entity

type Merchant struct {
	AcquirerBID           string `xml:"AcquirerBID"`
	AcquirerName          string `xml:"AcquirerName"`
	AcquirerMerchantID    string `xml:"AcquirerMerchantID"`
	CAIDs                 CAIDs
	AMMFID                string `xml:"AMMFID"` // @todo нужно уточнить так как всегда пустой
	DateSigned			  string `xml:"DateSigned"`
	MerchantAlias         MerchantAlias
	LocationAddress		  LocationAddress
	MCCs				  MCCs
	TaxDetail			  TaxDetail
	NonMonetaryChangeDate string `xml:"NonMonetaryChangeDate"`
	ChangeIndicator       string `xml:"ChangeIndicator,attr"`
	LocationCountry       string `xml:"LocationCountry,attr"`
}


func (merchant *Merchant) CreateMerchants(rowStruct map[string]string) *Merchant {
	merchant.AcquirerBID = rowStruct["ACQUIRERBID"]
	merchant.AcquirerName = rowStruct["ACQUIRERNAME"]
	merchant.AcquirerMerchantID = rowStruct["ACQUIRERMERCHANTID"]
	merchant.DateSigned = rowStruct["DATESIGNED"]
	merchant.AMMFID = ""
	merchant.NonMonetaryChangeDate = rowStruct["DATESIGNED"]


	merchant.CAIDs.CAID.CAID = rowStruct["CAID"]
	merchant.CAIDs.CAID.AcquirerBIN = rowStruct["ACQUIRERBIN"]

	merchant.MerchantAlias.DBAName = rowStruct["DBANAME"]
	merchant.MerchantAlias.LegalName = rowStruct["LEGALNAME"]
	merchant.MerchantAlias.CorporateName = rowStruct["CORPORATENAME"]
	merchant.MerchantAlias.BASEIIName = rowStruct["BASEIINAME"]



	merchant.LocationAddress.AddressToLine1 = rowStruct["ADDRESSTOLINE1"]
	merchant.LocationAddress.AddressToLine2 = ""
	merchant.LocationAddress.Street = rowStruct["STREET"]
	merchant.LocationAddress.City = rowStruct["CITY"]
	merchant.LocationAddress.StateProvinceCode = rowStruct["STATEPROVINCECODE"]
	merchant.LocationAddress.PostalCode = rowStruct["POSTALCODE"]



	merchant.MCCs.MCC.MCC = rowStruct["MCC"]
	merchant.MCCs.MCC.Sequence = "1"

	merchant.TaxDetail.CorporateStatus = rowStruct["CORPORATESTATUS"]

	return merchant
}