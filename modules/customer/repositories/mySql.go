package repository

import (
	"context"
	"database/sql"
	"log"

	. "github.com/gulliverwald/clean-arch-go/domain"
)

type MySqlRepository struct {
	Conn *sql.DB
}

func NewMySqlRepository() CustomerRepository {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/example")
	if err != nil {
		log.Fatal(err)
	}

	return &MySqlRepository{Conn: db}
}

func (mySql *MySqlRepository) Create(ctx context.Context, customer *Customer) error {
	query := `INSERT INTO customer (firstName, lastName, document) VALUES (@Firstname, @Lastname, @Document);`

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
