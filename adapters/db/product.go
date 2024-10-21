package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pedrojpx/hexagonal-example/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var prod application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Status)
	if err != nil {
		return nil, err
	}
	return &prod, nil
}

// func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {

// }
