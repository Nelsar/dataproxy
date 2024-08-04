package model

import (
	"time"
)

type Products struct {
	XMLName struct{} `xml:"products"`
	Product *Product `xml:"product"`
}

type Product struct {
	Id                      int               `xml:"id"`
	XMLName                 struct{}          `xml:"product"`
	ManufacturerId          int               `xml:"id_manufacturer"`
	SupplierId              int               `xml:"id_supplier"`
	DefaultCategoryId       int               `xml:"id_category_default"`
	New                     string            `xml:"new"`
	DefaultAttributeCache   string            `xml:"cache_default_attribute"`
	DefaultImageId          int               `xml:"id_default_image"`
	DefaultCombinationId    int               `xml:"id_default_combination"`
	TaxRulesGroupId         int               `xml:"id_tax_rules_group"`
	PositionInCategory      string            `xml:"position_in_category"`
	Type                    string            `xml:"type"`
	DefaultShopId           int               `xml:"id_shop_default"`
	Reference               string            `xml:"reference"`
	ReferenceSupplier       string            `xml:"supplier_reference"`
	Location                string            `xml:"location"`
	Width                   float32           `xml:"width"`
	Height                  float32           `xml:"height"`
	Depth                   float32           `xml:"depth"`
	Weight                  float32           `xml:"weight"`
	QuantityDiscount        int               `xml:"quantity_discount"`
	Ean13                   string            `xml:"ean13"`
	ISBN                    string            `xml:"isbn"`
	Upc                     string            `xml:"upc"`
	Mpn                     string            `xml:"mpn"`
	Quantity                int               `xml:"quantity"`
	CacheIsPack             int               `xml:"cache_is_pack"`
	CacheHasAttachments     int               `xml:"cache_has_attachments"`
	IsVirtual               bool              `xml:"is_virtual"`
	State                   int               `xml:"state"`
	AdditionalDeliveryTimes int               `xml:"additional_delivery_times"`
	DeliveryInStock         *DeliveryInStock  `xml:"delivery_in_stock"`
	DeliveryOutStock        *DeliveryOutStock `xml:"delivery_out_stock"`
	ProductType             string            `xml:"product_type"`
	OnSale                  int               `xml:"on_sale"`
	OnlineOnly              int               `xml:"online_only"`
	Ecotax                  float32           `xml:"ecotax"`
	MinimalQuantity         int               `xml:"minimal_quantity"`
	LowStockThreshold       int               `xml:"low_stock_threshold"`
	LowStockAlert           int               `xml:"low_stock_alert"`
	Price                   float32           `xml:"price"`
	WholesalePrice          float32           `xml:"wholesale_price"`
	Unity                   string            `xml:"unity"`
	UnitPriceRatio          float32           `xml:"unit_price_ratio"`
	AdditionalShippingCost  float32           `xml:"additional_shipping_cost"`
	Customizable            int               `xml:"customizable"`
	TextFields              string            `xml:"text_fields"`
	UploadableFiles         string            `xml:"uploadable_files"`
	Active                  int               `xml:"active"`
	RedirectType            string            `xml:"redirect_type"`
	RedirectedTypeId        int               `xml:"redirected"`
	AvailableForOrder       int               `xml:"available_for_order"`
	AvailableDate           time.Time         `xml:"available_date"`
	ShowCondition           int               `xml:"show_condition"`
	Condition               string            `xml:"condition"`
	ShowPrice               int               `xml:"show_price"`
	Indexed                 int               `xml:"indexed"`
	Visibility              int               `xml:"visibility"`
	AdvancedStockManagement string            `xml:"advanced_stock_management"`
	DateAdd                 time.Time         `xml:"date_add"`
	DateUpd                 time.Time         `xml:"date_upd"`
	PackStockType           int               `xml:"pack_stock_type"`
	MetaDescription         *MetaDescription  `xml:"meta_description"`
	MetaKeywords            *MetaKeywords     `xml:"meta_keywords"`
	MetaTitle               *MetaTitle        `xml:"meta_title"`
	LinkRewrite             *LinkRewrite      `xml:"link_rewrite"`
	Name                    *Name             `xml:"name"`
	Description             *Description      `xml:"description"`
	DescriptionShort        *DescriptionShort `xml:"short_description"`
	AvailableNow            *AvailableNow     `xml:"available_now"`
	AvailableLater          *AvailableLater   `xml:"available_later"`
	ProductAttributeId      int               `xml:"id_product_attribute"`
	Associations            *Associations     `xml:"associations"`
}

type Language struct {
	Value   string   `xml:",chardata"`
	Id      int      `xml:"id"`
	XMLName struct{} `xml:"language"`
}

type DeliveryInStock struct {
	XMLName  struct{}  `xml:"delivery_in_stock"`
	Language *Language `xml:"language"`
}

type DeliveryOutStock struct {
	XMLName  struct{}  `xml:"delivery_out_stock"`
	Language *Language `xml:"language"`
}

type MetaDescription struct {
	XMLName  struct{}  `xml:"meta_description"`
	Language *Language `xml:"language"`
}

type MetaKeywords struct {
	XMLName  struct{}  `xml:"meta_keywords"`
	Language *Language `xml:"language"`
}

type MetaTitle struct {
	XMLName  struct{}  `xml:"meta_title"`
	Language *Language `xml:"language"`
}

type LinkRewrite struct {
	XMLName  struct{}  `xml:"link_rewrite"`
	Language *Language `xml:"language"`
}

type Name struct {
	XMLName  struct{}  `xml:"name"`
	Language *Language `xml:"language"`
}

type Description struct {
	XMLName  struct{}  `xml:"description"`
	Language *Language `xml:"language"`
}

type DescriptionShort struct {
	XMLName  struct{}  `xml:"short_description"`
	Language *Language `xml:"language"`
}

type AvailableNow struct {
	XMLName  struct{}  `xml:"available_now"`
	Language *Language `xml:"language"`
}

type AvailableLater struct {
	XMLName  struct{}  `xml:"available_later"`
	Language *Language `xml:"language"`
}

type ProductOptionValues struct {
	XMLName struct{} `xml:"product_option_values"`
}

type ProductOptionValue struct {
	XMLName struct{} `xml:"product_option_value"`
	Id      int      `xml:"id"`
}

type ProductFeatures struct {
	XMLName struct{} `xml:"product_features"`
}

type ProductFeature struct {
	XMLName        struct{} `xml:"product_feature"`
	Id             int      `xml:"id"`
	FeatureValueId int      `xml:"id_feature_value"`
}

type StockAvailables struct {
	XMLName        struct{}        `xml:"stock_availables"`
	StockAvailable *StockAvailable `xml:"stock_available"`
}

type StockAvailable struct {
	XMLName            struct{} `xml:"stock_available"`
	Id                 int      `xml:"id"`
	ProductAttributeId int      `xml:"id_product_attribute"`
}

type Accessories struct {
	XMLName struct{} `xml:"accessories"`
	Product *Product `xml:"product"`
}

type ProductBundle struct {
	XMLName struct{} `xml:"product_bundle"`
	Product *Product `xml:"product"`
}
