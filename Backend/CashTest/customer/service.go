package customer

import (
	mgo "gopkg.in/mgo.v2"
)

type CustomerService interface {
	FindAll() ([]*Customer, error)
	FindByAccountNumber(accountNumber string) (*Customer, error)
	Update(accountNumber string, customer Customer) (*mgo.ChangeInfo, error)
}
