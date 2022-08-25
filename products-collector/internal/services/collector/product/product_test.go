package product

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
	"testing"
)

func TestProduct(t *testing.T) {
	xml := getProductXML()
	product := ParseProductXML(xml)
	images := product.Images
	details := product.File.Data
	expected := "http://cdn.cnetcontent.com/18/e4/18e4d00c-bfbf-4fb4-bb3b-58267f4b722a.jpg"
	actual := images[0].Data

	if len(images) != 2 {
		t.Fatalf("images size does not match")
	}

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected %s\nactual: %s", expected, actual)
	}
	if strings.Compare("http://website.com/datasheet.do", details) != 0 {
		t.Fatalf("details url does not match")
	}
}

func getProductXML() string {
	content, err := ioutil.ReadFile("product_test.xml")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
