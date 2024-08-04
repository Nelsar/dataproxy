package db

import (
	"database/sql"
	"dataproxy/model"
	"log"
)

func GetElements() []model.Elements {
	db := DbConnect()

	rows, err := db.Query("select `Id`, `Name`, `fullname`, `quantity`, `article_pn`, `article`, `price`, `brand`, `weight`, `category` from Elements")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	elements := []model.Elements{}

	for rows.Next() {
		e := model.Elements{}
		err := rows.Scan(
			&e.Id,
			&e.Name,
			&e.FullName,
			&e.Quantity,
			&e.ArticlePn,
			&e.Article,
			&e.Price,
			&e.Brand,
			&e.Weight,
			&e.CategoryId)
		if err != nil {
			log.Panicln()
			continue
		}
		elements = append(elements, e)
	}
	defer rows.Close()
	return elements
}

func GetElement(article string) (model.Elements, error) {
	db := DbConnect()

	element := model.Elements{}

	row := db.QueryRow("select `Id`, `Name`, `fullname`, `quantity`, `article`, `article_pn`, `price`, `brand`, `weight`, `category` from Elements where `article` = ?", article)

	err := row.Scan(
		&element.Id,
		&element.Name,
		&element.FullName,
		&element.Quantity,
		&element.Article,
		&element.ArticlePn,
		&element.Price,
		&element.Brand,
		&element.Weight,
		&element.CategoryId)

	defer db.Close()

	return element, err
}

func UpdateElement(element model.Elements) sql.Result {
	db := DbConnect()
	articlePn, err := element.ArticlePn.Value()
	if err != nil {
		log.Fatalln(err)
	}

	brand, err := element.Brand.Value()
	if err != nil {
		log.Fatalln(err)
	}

	weight, err := element.Weight.Value()
	if err != nil {
		log.Fatalln(err)
	}

	result, err := db.Exec("update Elements set `Name` = ?, `fullname` = ?, `quantity` = ?, `article_pn` = ?, `price` = ?, `brand` = ?, `weight` = ?,  `category` = ? where `article` = ?", &element.Name,
		&element.FullName,
		&element.Quantity,
		&articlePn,
		&element.Price,
		&brand,
		&weight,
		&element.CategoryId,
		&element.Article)

	log.Printf("Updated Article: %s", articlePn)

	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	return result

}

func AddElement(element model.Elements) sql.Result {
	db := DbConnect()

	result, err := db.Exec("insert into Elements (`Name`, `fullname`, `quantity`, `article`, `article_pn`, `price`, `brand`, `weight`, `category`) values(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		&element.Name,
		&element.FullName,
		&element.Quantity,
		&element.Article,
		&element.ArticlePn,
		&element.Price,
		&element.Brand,
		&element.Weight,
		&element.CategoryId)

	log.Printf("Added Article: %s", element.ArticlePn)

	if err != nil {
		log.Fatalln(err)
	}
	db.Close()

	return result
}

func GetElementsFullDelete() sql.Result {
	db := DbConnect()

	result, err := db.Exec("TRUNCATE TABLE `Elements`")
	if err != nil {
		log.Fatalln()
	}
	db.Close()

	return result
}
