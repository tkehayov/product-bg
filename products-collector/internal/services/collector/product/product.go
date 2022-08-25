package product

import (
	"encoding/xml"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Product struct {
	File   File    `xml:"file"`
	Images []Image `xml:"image"`
}
type File struct {
	Data string `xml:"data"`
}

type Image struct {
	Data string `xml:"data"`
}

func ParseProductXML(content string) Product {
	var product Product
	content = adjustXml(content)
	err := xml.Unmarshal([]byte(content), &product)

	if err != nil {
		log.Error("error", err)
	}

	return product
}

func adjustXml(content string) string {
	content = strings.ReplaceAll(content, "<![CDATA[ ", "<data>")
	content = strings.ReplaceAll(content, " ]]>", "</data>")

	return content
}
