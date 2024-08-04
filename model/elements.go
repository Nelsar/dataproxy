package model

import (
	"database/sql"
)

type Elements struct {
	Id          int
	Name        string
	FullName    string
	Quantity    sql.NullString
	ArticlePn   sql.NullString
	Article     string
	Price       float64
	Price2      float64
	Brand       sql.NullString
	Weight      sql.NullString
	CategoryId  int
	Description string
	Warranty    string
}

type ElementsRoot struct {
	Elements   []ElementsViewModel `json:"elements"`
	Pagination Pagination          `json:"pagination"`
}

type ElementsViewModel struct {
	Id          int
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	Quantity    interface{} `json:"quantity"`
	ArticlePn   string      `json:"article_pn"`
	Article     int         `json:"article"`
	Price       float64     `json:"price1"`
	Price2      float64     `json:"price2"`
	Brand       string      `json:"brand"`
	Weight      float64     `json:"weight"`
	CategoryId  int         `json:"category"`
	Description string      `json:"description"`
	Warranty    string      `json:"warranty"`
}

type Pagination struct {
	TotalCount  int `json:"totalCount"`
	TotalPages  int `json:"totalPages"`
	CurrentPage int `json:"currentPage"`
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
}
