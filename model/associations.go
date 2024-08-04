package model

type Associations struct {
	Categories          Categories          `xml:"categories"`
	Images              Images              `xml:"images"`
	Combinations        Combinations        `xml:"combinations"`
	ProductOptionValues ProductOptionValues `xml:"product_option_values"`
	ProductFeatures     ProductFeatures     `xml:"product_features"`
	Tags                Tags                `xml:"tags"`
	StockAvailables     StockAvailables     `xml:"stock_availables"`
	Attachments         Attachments         `xml:"attachments"`
	Accessories         Accessories         `xml:"accessories"`
	ProductBundle       ProductBundle       `xml:"product_bundle"`
	Addresses           Addresses           `xml:"addresses"`
	Products            Products            `xml:"products"`
}
