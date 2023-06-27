package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/gulliverwald/clean-arch-go/domain"
)

type MySqlRepository struct {
	Conn *sql.DB
}

func NewMySqlRepository() ServiceRepository {
	db, err := sql.Open("mysql", "user:password@tcp(0.0.0.0:3306)/example")
	if err != nil {
		log.Fatal(err)
	}

	return &MySqlRepository{Conn: db}
}

func (mySql *MySqlRepository) Create(ctx context.Context, service *Service) error {
	res, err := mySql.Conn.Exec(
		`INSERT INTO service (name, price, duration, scheduleDate, customerId) VALUES (?, ?, ?, ?, ?)`,
		service.Name,
		service.Price,
		service.Duration,
		service.ScheduleDate,
		service.CustomerID,
	)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1452 (23000)") {
			return errors.New("Customer does not exists.")
		}

		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	service.ID = lastID

	return nil
}

func (mySql *MySqlRepository) Fetch(ctx context.Context) ([]Service, error) {
	services := []Service{}

	rows, err := mySql.Conn.Query(`SELECT id, name, price, duration, scheduleDate, customerId, createdAt, updatedAt	FROM service`)
	if err != nil {
		return services, err
	}

	defer rows.Close()

	for rows.Next() {
		var service Service
		err = rows.Scan(
			&service.ID,
			&service.Name,
			&service.Price,
			&service.Duration,
			&service.ScheduleDate,
			&service.CustomerID,
			&service.CreatedAt,
			&service.UpdatedAt,
		)
		if err != nil {
			return services, err
		}

		services = append(services, service)
	}

	return services, nil
}

func (mySql *MySqlRepository) GetByID(ctx context.Context, id int64) (Service, error) {
	service := Service{}

	err := mySql.Conn.QueryRow(`SELECT * FROM service WHERE id = ?`, id).Scan(
		&service.ID,
		&service.Name,
		&service.Price,
		&service.Duration,
		&service.ScheduleDate,
		&service.CustomerID,
		&service.CreatedAt,
		&service.UpdatedAt,
	)
	if err != nil {
		return Service{}, err
	}

	return service, nil
}

func (mySql *MySqlRepository) GetServiceByScheduleDateAndCustomerId(
	ctx context.Context,
	customerId string,
	scheduleDate string,
) (Service, error) {
	service := Service{}

	err := mySql.Conn.QueryRow(
		`SELECT 
			id,
			name,
			price,
			duration,
			scheduleDate,
			customerId,
			createdAt,
			updatedAt
		FROM service
		WHERE 
			customerId = ?
			AND scheduleDate = ?`,
		customerId,
		scheduleDate,
	).Scan(
		&service.ID,
		&service.Name,
		&service.Price,
		&service.Duration,
		&service.ScheduleDate,
		&service.CustomerID,
		&service.CreatedAt,
		&service.UpdatedAt,
	)
	if err != nil {
		return Service{}, err
	}

	return service, nil
}

func (mySql *MySqlRepository) Update(ctx context.Context, service *Service) error {
	_, err := mySql.Conn.Exec(
		`UPDATE service SET name = ?, price = ?, duration = ?, scheduleDate = ? WHERE id = ?`,
		service.Name,
		service.Price,
		service.Duration,
		service.ScheduleDate,
		service.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (mySql *MySqlRepository) Delete(ctx context.Context, id int64) error {
	_, err := mySql.Conn.Exec(`DELETE FROM service WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
