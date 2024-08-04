package service

import (
	"dataproxy/configuration"
	"dataproxy/db"
	"dataproxy/model"
	"log"
	"sync"
)

func CurrencyService(wg *sync.WaitGroup) {
	config := configuration.Exec().Currency
	response := GetData(config.Url)

	var currencies model.Currencies
	XmlDeserialze(response, &currencies)

	for _, value := range currencies.Valute {

		currency := model.Currency{
			Currency: value.Code,
			Rate:     value.Nominal / value.UnitRate,
		}

		exists := db.GetCurrency(currency.Currency)

		if len(exists.Currency) > 0 {
			rows := db.UpdateCurrency(currency)
			_, err := rows.RowsAffected()
			if err != nil {
				log.Fatalln(err)
			}
		}

		if len(exists.Currency) == 0 {
			rows := db.AddCurrency(currency)
			_, err := rows.RowsAffected()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	wg.Done()
}
