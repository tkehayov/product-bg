package product_list

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
	"testing"
)

func TestProductList(t *testing.T) {
	xml := getProductsListXML()
	productList := ParseProductListXML(xml)
	actualCodeID := productList.Products[1].CodeId
	expectedCodeID := "MVVK2ZE/A"

	if len(productList.Products) != 2 {
		t.Fatalf("product size does not match")
	}

	if strings.Compare(expectedCodeID, actualCodeID) != 0 {
		t.Fatalf("codeId does not match")
	}
}

func getProductsListXML() string {
	content, err := ioutil.ReadFile("product_list_test.xml")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
