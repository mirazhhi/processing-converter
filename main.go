package main

import (
	"errors"
	"fmt"
	"web/entity"
	"encoding/xml"
	"io/ioutil"
	"time"
	"os"
	"log"
	"strings"
	"github.com/tealeg/xlsx/v3"
	"github.com/joho/godotenv"
)

var collect []string
var maxRow int




type Merchants struct {
	Merchant []entity.Merchant `xml:"Merchant"`
	Count    int `xml:"Count,attr"`
}

type AMMF struct {
	Merchants Merchants
	CreateDate      string `xml:"CreateDate,attr"`
    ProcessorBINCIB string `xml:"ProcessorBINCIB,attr"`
    ProcessorName   string `xml:"ProcessorName,attr"`
    FileSequence    string `xml:"FileSequence,attr"`
    Version         string `xml:"Version,attr"`
}

func (ammf *AMMF) CreateStructure() {
	dataTime , _ := time.Now().UTC().MarshalText()
	dataTimeString := string(dataTime)

	ammf.CreateDate = strings.Split(dataTimeString, ".")[0]
	ammf.ProcessorBINCIB = os.Getenv("PROCESSOR_BINCIB")
	ammf.ProcessorName = os.Getenv("PROCESSOR_NAME")
	ammf.FileSequence = os.Getenv("FILE_SEQUENCE")
	ammf.Version = os.Getenv("VERSION")

	fmt.Println("Struct Created")
}

func cellVisitor(c *xlsx.Cell) error {
    value, err := c.FormattedValue()

    if err != nil {
        fmt.Println(err.Error())
    } else {
        // fmt.Println("Cell value:", value)

        collect = append(collect, value)
    }
    return err
}

func rowVisitor(r *xlsx.Row) error {
    return r.ForEachCell(cellVisitor)
}

func rowStuff() {

    filename := os.Getenv("XLSX_PATH")
    wb, err := xlsx.OpenFile(filename)
    if err != nil {
        panic(err)
    }
    sh, ok := wb.Sheet[os.Getenv("SHHET_NAME")]
    if !ok {
        panic(errors.New("Sheet not found"))
    }
    sh.ForEachRow(rowVisitor)
    maxRow = sh.MaxRow
    fmt.Println("Max row is", sh.MaxRow)

    cellRow := len(collect) / sh.MaxRow
    fmt.Println("Max cellrow is", cellRow)
    
}

func createCollect(ammf *AMMF) {
	merch := entity.Merchant{}
	merchCollect := []entity.Merchant{}
	cellCountInOneRow := len(collect) / maxRow

	rowStruct := map[string]string{}

	for i := 1; i < maxRow; i++ {
		rowCollect := collect[(cellCountInOneRow * i):(cellCountInOneRow * i) + cellCountInOneRow]
		
		rowStruct["PROCESSORBINCIB"] = rowCollect[0]
		rowStruct["PROCESSORNAME"] = rowCollect[1]
		rowStruct["LOCATIONCOUNTRY"] = rowCollect[2]
		rowStruct["ACQUIRERBID"] = rowCollect[3]
		rowStruct["ACQUIRERNAME"] = rowCollect[4]
		rowStruct["ACQUIRERMERCHANTID"] = rowCollect[5]
		rowStruct["ACQUIRERBIN"] = rowCollect[6]
		rowStruct["CHANGEINDICATOR"] = rowCollect[7]
		rowStruct["CAID"] = rowCollect[8]
		rowStruct["DATESIGNED"] = rowCollect[9]
		rowStruct["DBANAME"] = rowCollect[10]
		rowStruct["LEGALNAME"] = rowCollect[11]
		rowStruct["CORPORATENAME"] = rowCollect[12]
		rowStruct["BASEIINAME"] = rowCollect[13]
		rowStruct["ADDRESSTOLINE1"] = rowCollect[14]
		rowStruct["STREET"] = rowCollect[15]
		rowStruct["CITY"] = rowCollect[16]
		rowStruct["STATEPROVINCECODE"] = rowCollect[17]
		rowStruct["POSTALCODE"] = rowCollect[18]
		rowStruct["MCC"] = rowCollect[19]
		rowStruct["CORPORATESTATUS"] = rowCollect[20]

		merch.CreateMerchants(rowStruct)
		merchCollect = append(merchCollect, merch)
	}

    ammf.Merchants.Count = len(merchCollect)
    ammf.Merchants.Merchant = merchCollect
	fmt.Println(len(merchCollect))
}

func main () {
	
	fmt.Println("== CREATE XML FROM XLSX ==")
	err := godotenv.Load()
	if err != nil {
	  	log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("SHHET_NAME"))

    rowStuff()

    ammf := AMMF{}

    ammf.CreateStructure()

    createCollect(&ammf)

 
	file, _ := xml.MarshalIndent(ammf, "", "\t")
 
	_ = ioutil.WriteFile(os.Getenv("XML_FILE_NAME"), file, 0644)
}