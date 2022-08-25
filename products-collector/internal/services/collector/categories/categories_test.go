package categories

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCategoriesList(t *testing.T) {
	xml := getCategoriesXML()
	categories := ParseCategoriesXML(xml)
	productCategory := categories.ProductCategory[1]

	productCategoryNotMatch, messageCategory := productCategoryNotMatch(categories.ProductCategory)
	productGroupNotMatch, messageGroup := productGroupNotMatch(productCategory.ProductGroup)
	linkNotMatch, messageLink := productLinkNotMatch(productCategory.ProductGroup[0])

	if productCategoryNotMatch {
		t.Fatalf(messageCategory)
	}

	if productGroupNotMatch {
		t.Fatalf(messageGroup)
	}

	if productGroupNotMatch {
		t.Fatalf(messageGroup)
	}

	if linkNotMatch {
		t.Fatalf(messageLink)
	}
}

func productLinkNotMatch(group ProductGroup) (bool, string) {
	actual := group.AtomLink.URL
	expected := "https://b2b.also.com/invoke/ActDelivery_HTTP.Inbound/receiveXML_API?j_u=11134342&j_p=E0r34A6d_s_D46B&propertyId=B02001004"
	if strings.Compare(actual, expected) != 0 {
		return true, fmt.Sprintf("expected: %s\n actual: %s", expected, actual)
	}

	return false, ""
}

func productCategoryNotMatch(productCategory []ProductCategory) (bool, string) {
	expected := "Дисплеи, монитори и проектори"
	actual := productCategory[0].Name

	if len(productCategory) != 2 {
		return true, "product category size is not 2"
	}

	if strings.Compare(actual, expected) != 0 {
		return true, fmt.Sprintf("expected: %s\n actual: %s", expected, actual)
	}

	return false, ""
}

func productGroupNotMatch(group []ProductGroup) (bool, string) {
	expected := "Аксесоари за телевизори"
	actual := group[0].PropertyGroupId

	if len(actual) != 2 {
		return true, "product group size is not 2"
	}

	if strings.Compare(actual[1], expected) != 0 {
		return true, fmt.Sprintf("expected: %s\n actual: %s", expected, actual)
	}

	return false, ""
}

func getCategoriesXML() string {
	content, err := ioutil.ReadFile("categories_test.xml")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
