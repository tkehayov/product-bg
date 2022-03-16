package products

type ProductDto struct {
	MerchantId string     `xml:"merchantId" json:"merchantId"`
	Products   []*Product `xml:"products" json:"products"`
}

type Product struct {
	CodeId       string  `xml:"codeId" json:"codeId"`
	Price        float64 `xml:"price" json:"price"`
	ShippingFee  float64 `xml:"shippingFee" json:"shippingFee"`
	ProductTitle string  `xml:"productTitle" json:"productTitle"`
	Url          string  `xml:"url" json:"url"`
}
