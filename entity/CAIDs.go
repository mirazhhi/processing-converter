package entity



type CAIDs struct {
	CAID CAID `xml:"CAID"`
}


type CAID struct {
	CAID string `xml:",chardata"`
	AcquirerBIN string `xml:"AcquirerBIN,attr"`
}