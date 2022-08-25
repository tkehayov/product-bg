package categories

import (
	"encoding/xml"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type Catalog struct {
	ProductCategory []ProductCategory `xml:"productCategory"`
}

type ProductCategory struct {
	Name         string         `xml:"name,attr"`
	ProductGroup []ProductGroup `xml:"productGroup"`
}

type ProductGroup struct {
	PropertyGroupId []string `xml:"propertyGroupId"`
	AtomLink        AtomLink `xml:"atom"`
}

type AtomLink struct {
	URL string `xml:"href,attr"`
}

func ParseCategoriesXML(content string) Catalog {
	var catalog Catalog
	content = adjustXml(content)
	d := xml.NewDecoder(os.Stdin)
	d.Strict = false
	err := xml.Unmarshal([]byte(content), &catalog)
	if err != nil {
		log.Error("errror", err)
	}

	return catalog
}

func adjustXml(content string) string {
	content = strings.ReplaceAll(content, "&", "&amp;")
	content = strings.ReplaceAll(content, "atom:link", "atom")

	return content
}
