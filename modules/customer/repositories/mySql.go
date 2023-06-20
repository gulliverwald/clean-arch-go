package repository

import (
	"context"
	"log"
	"database/sql"

	. "github.com/gulliverwald/clean-arch-go/domain"
)

type MySqlRepository struct {
	Conn *sql.DB
}

func NewMySqlRepository() CustomerRepository {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/example")

	if err != nil {
		log.Fatal(err)
	}

	return &MySqlRepository{ db }
}

func (mySql *MySqlRepository) Create(ctx context.Context, customer *Customer) error {
	query := `INSERT INTO customer (firstName, lastName, document) VALUES firstName = $1, lastName = $2, document = $3`

	stmt, err := mySql.Conn.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, customer.Firstname, customer.Lastname, customer.Document)

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
	return []Customer{}, nil
}

func (mySql *MySqlRepository) GetByID(ctx context.Context, id int64) (Customer, error) {
	return Customer{}, nil
}

func (mySql *MySqlRepository) Update(ctx context.Context, customer *Customer) error {
	return nil
}

func (mySql *MySqlRepository) Delete(ctx context.Context, id int64) error {
	return nil
}