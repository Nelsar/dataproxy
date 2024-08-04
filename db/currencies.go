package db

import (
	"database/sql"
	"dataproxy/model"
	"log"
)

func GetCurrency(code string) model.Currency {
	db := DbConnect()

	row := db.QueryRow("select `Id`, `Currency`, `Rate` from `Currencies` where `Currency` = ?", code)

	currency := model.Currency{}
	err := row.Scan(&currency.Id, &currency.Currency, &currency.Rate)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalln(err)
		}
		log.Fatalln(err)
	}
	defer db.Close()

	return currency
}

func AddCurrency(currency model.Currency) sql.Result {
	db := DbConnect()

	result, err := db.Exec("INSERT INTO `Currency`(`Currency`, `Rate`) VALUES(?, ?)", &currency.Currency, &currency.Rate)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	return result

}

func UpdateCurrency(currency model.Currency) sql.Result {
	db := DbConnect()

	result, err := db.Exec("UPDATE `Currency` SET `Rate` = ? WHERE `Currency` = ?", &currency.Rate, &currency.Currency)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	return result
}
