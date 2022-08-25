package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func main() {
	//https://b2b.also.com/invoke/ActDelivery_HTTP.Inbound/receiveXML_API?CatalogCategory=true&j_u=11134342&j_p=E0r34A6d_s_D46B
	//https://b2b.also.com/invoke/ActDelivery_HTTP.Inbound/receiveXML_API?j_u={{username}}&j_p={{password}}&ProductData=true&codeId=4015617
	url := getUrl()
	log.Error(url)
}

//https://b2b.also.com/invoke/ActDelivery_HTTP.Inbound/receiveXML_API?j_u=11134342&j_p=E0r34A6d_s_D46B&ProductData=true&codeId=4015617
//https://b2b.also.com/invoke/ActDelivery_HTTP.Inbound/receiveXML_API?j_u=USERNAME&j_p=PASSWORD&ProductData=true&codeId=4015617
func getUrl() string {
	alsoBaseUrl := os.Getenv("ALSO_BASE_URL")
	//categoryUrl := alsoBaseUrl + "&CatalogCategory=true"
	//productsUrl := alsoBaseUrl + "&ProductData=true&codeId=" + 4015617
	return alsoBaseUrl
}
