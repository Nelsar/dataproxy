package service

import (
	"database/sql"
	"dataproxy/configuration"
	"dataproxy/db"
	"dataproxy/model"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

func XmlDeserialze(data []byte, v any) {
	err := xml.Unmarshal(data, v)
	if err != nil {
		log.Fatal(err)
	}
}

func XmlDataAddToDatabase() {
	byteValue := GetDownload(configuration.Exec().Distributor.Yml)

	var catalog model.Catalog
	XmlDeserialze(byteValue, &catalog)

	offer := catalog.Shop.Offers.Offer

	for i := 0; i < len(offer); i++ {

		var e model.Elements
		e.Name = offer[i].Name
		e.FullName = offer[i].Name
		e.Quantity.String = offer[i].QuantityInStock
		e.Quantity.Valid = true
		e.Article = strconv.Itoa(offer[i].VendorCode)

		for p := 0; p < len(offer[i].Param); p++ {
			if offer[i].Param[p].Name == "Артикул" {
				e.ArticlePn.String = offer[i].Param[p].Value
				e.ArticlePn.Valid = true
			}

			if offer[i].Param[p].Name == "Вес" {
				e.Weight.String = offer[i].Param[p].Value
				e.Weight.Valid = true
			}
		}

		e.Price = offer[i].PurchasePrice
		e.Brand.String = offer[i].Vendor
		e.Brand.Valid = true

		exist, err := db.GetElement(strconv.Itoa(offer[i].VendorCode))

		if err != nil {
			if err != sql.ErrNoRows {
				log.Fatalln(err)
			}

			if err == sql.ErrNoRows {
				ir := db.AddElement(e)
				row, ierr := ir.RowsAffected()
				if ierr != nil {
					log.Fatalln(ierr)
				}
				fmt.Printf("Xml Added Rows: %d, Article: %s\n", row, exist.Article)
			}
		}

		if len(exist.Article) > 0 {
			ur := db.UpdateElement(e)
			urow, uerr := ur.RowsAffected()
			if uerr != nil {
				log.Fatalln(uerr)
			}
			fmt.Printf("Xml Updated Rows: %d, Article: %s\n", urow, exist.Article)
		}
	}
}
