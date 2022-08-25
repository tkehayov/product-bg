package product_list

import (
	"encoding/xml"
	log "github.com/sirupsen/logrus"
)

type ProductSet struct {
	Products []Product `xml:"product"`
	//todo check if vendor is needed here
}

type Product struct {
	CodeId string `xml:"codeId,attr"`
}

func ParseProductListXML(content string) ProductSet {
	var productSet ProductSet
	err := xml.Unmarshal([]byte(content), &productSet)

	if err != nil {
		log.Error("errror", err)
	}

	return productSet
}
