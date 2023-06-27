package repository

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/gulliverwald/clean-arch-go/domain"
)

type MySqlRepository struct {
	Conn *sql.DB
}

func NewMySqlRepository() CustomerRepository {
	db, err := sql.Open("mysql", "user:password@tcp(0.0.0.0:3306)/example")
	if err != nil {
		log.Fatal(err)
	}

	return &MySqlRepository{Conn: db}
}

func (mySql *MySqlRepository) Create(ctx context.Context, customer *Customer) error {
	res, err := mySql.Conn.Exec(
		`INSERT INTO customer (firstName, lastName, document) VALUES (?, ?, ?)`,
		customer.Firstname,
		customer.Lastname,
		customer.Document,
	)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	customer.ID = lastID

	return nil
}

func (mySql *MySqlRepository) Fetch(ctx context.Context) ([]Customer, error) {
	customers := []Customer{}

	rows, err := mySql.Conn.Query(`SELECT * FROM customer`)
	if err != nil {
		return customers, err
	}

	defer rows.Close()

	for rows.Next() {
		var customer Customer
		err = rows.Scan(
			&customer.ID,
			&customer.Firstname,
			&customer.Lastname,
			&customer.Document,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)
		if err != nil {
			return customers, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (mySql *MySqlRepository) GetByID(ctx context.Context, id int64) (Customer, error) {
	customer := Customer{}

	err := mySql.Conn.QueryRow(`SELECT * FROM customer WHERE id = ?`, id).Scan(
		&customer.ID,
		&customer.Firstname,
		&customer.Lastname,
		&customer.Document,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		return Customer{}, err
	}

	return customer, nil
}

func (mySql *MySqlRepository) Update(ctx context.Context, customer *Customer) error {
	_, err := mySql.Conn.Exec(
		`UPDATE customer SET firstName = ?, lastName = ?, document = ? WHERE id = ?`,
		customer.Firstname,
		customer.Lastname,
		customer.Document,
		customer.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (mySql *MySqlRepository) Delete(ctx context.Context, id int64) error {
	_, err := mySql.Conn.Exec(`DELETE FROM customer WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
