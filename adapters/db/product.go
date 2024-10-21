package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	app "github.com/pedrojpx/hexagonal-example/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (app.ProductInterface, error) {
	var prod app.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Status)
	if err != nil {
		return nil, err
	}
	return &prod, nil
}

func (p *ProductDb) Save(product app.ProductInterface) (app.ProductInterface, error) {
	var rows int
	p.db.QueryRow("select count(*) from products where id = ?", product.GetId()).Scan(&rows)

	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDb) create(prod app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare("insert into products(id, name, price, status) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(prod.GetId(), prod.GetName(), prod.GetPrice(), prod.GetStatus())
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func (p *ProductDb) update(prod app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare("update products set name=?, price=?, status=? where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(prod.GetName(), prod.GetPrice(), prod.GetStatus(), prod.GetId())
	if err != nil {
		return nil, err
	}
	return prod, nil
}
