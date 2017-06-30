package dao

import (
	"summary/domain"
	"fmt"
	"database/sql"
)

type DbHandler interface {
	Execute(statement string)
	Query(statement string) *sql.Rows
}

type DbRepo struct {
	dbHandler DbHandler
}

type CustomerDbRepo DbRepo

func NewCustomerRepo(handler DbHandler) *CustomerDbRepo {
	repo := new(CustomerDbRepo)
	repo.dbHandler = handler
	return repo
}

func (repo *CustomerDbRepo) AddCustomer(customer domain.Customer) {
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO customers (id, name)
                                        VALUES ('%d', '%v')`,
		customer.Id, customer.Name))
}

func (repo *CustomerDbRepo) FindById(id int) domain.Customer {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT name FROM customers WHERE id = '%d' LIMIT 1", id))
	var name string
	row.Next()
	row.Scan(&name)
	return domain.Customer{Id: id, Name: name}
}
