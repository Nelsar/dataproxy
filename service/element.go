package service

import (
	"database/sql"
	"dataproxy/configuration"
	"dataproxy/db"
	"dataproxy/model"
	"fmt"
	"log"
	"strconv"
	"time"
)

func ElementsUpsert() {
	config := configuration.Exec().Distributor
	url := config.Api

	i := 1
	var getElements model.ElementsRoot

	for {
		time.Sleep(5 * time.Second)
		getElements = GetJsonData[model.ElementsRoot](fmt.Sprintf("%s&offset=%d", url, i))

		for _, element := range getElements.Elements {

			exists, err := db.GetElement(fmt.Sprintf("%d", element.Article))

			exists.Name = element.Name
			exists.FullName = element.FullName
			exists.Quantity.String = fmt.Sprintf("%v", element.Quantity)
			exists.Quantity.Valid = true
			exists.Article = strconv.Itoa(element.Article)
			exists.ArticlePn.String = element.ArticlePn
			exists.ArticlePn.Valid = true
			exists.Brand.String = element.Brand
			exists.Brand.Valid = true
			exists.Price = element.Price
			exists.Price2 = element.Price2
			exists.Weight.String = fmt.Sprintf("%f", element.Weight)
			exists.CategoryId = element.CategoryId
			exists.Warranty = element.Warranty

			fmt.Println(element.Price)
			if err != nil {
				if err != sql.ErrNoRows {
					log.Fatalln(err)
				}
				if err == sql.ErrNoRows {
					ir := db.AddElement(exists)
					_, ierr := ir.RowsAffected()
					if ierr != nil {
						log.Fatalln(ierr)
					}
				}
			}

			if len(exists.Article) > 0 {

				ur := db.UpdateElement(exists)
				_, uerr := ur.RowsAffected()
				if uerr != nil {
					log.Fatalln(uerr)
				}
			}
		}

		i = getElements.Pagination.Limit + getElements.Pagination.Offset + 1

		if len(getElements.Elements) == 0 {
			break
		}

	}
}
