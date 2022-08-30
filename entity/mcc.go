package entity


type MCCs struct {
	MCC MCC `xml:"MCC"`
}


type MCC struct {
	MCC string `xml:",chardata"`
	Sequence string `xml:"Sequence,attr"`
}