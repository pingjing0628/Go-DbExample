package models

import (
	"database/sql"
	"GoMySQL/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product {
					Id:       id, 
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Search(keyword string) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", "%" + keyword + "%")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product {
					Id:       id, 
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) SearchPrices(min, max float64) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where price >= ? and price <= ?", min, max)
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product {
					Id:       id, 
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}