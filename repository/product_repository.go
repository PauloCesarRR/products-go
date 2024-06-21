package repository

import (
	"database/sql"
	"fmt"
	"products/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}
	var products []model.Product
	var product model.Product

	for rows.Next() {
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		products = append(products, product)
	}

	rows.Close()

	return products, nil
}

func (pr *ProductRepository) GetProductByID(id int) (model.Product, error) {
	testr := "SELECT * FROM product WHERE id=%1"
	query, err := pr.connection.Prepare(testr)

	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, nil
		}
		fmt.Println(err)
		return model.Product{}, err
	}

	query.Close()
	return product, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	sql := "INSERT INTO product (name_product, price) VALUES ($1, $2) RETURNING id"
	query, err := pr.connection.Prepare(sql)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
