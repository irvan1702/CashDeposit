package mongo

import (
	"CashTest/customer"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CustomerService struct {
	collection *mgo.Collection
}

func NewCustomerService(session *Session, dbName string, collectionName string) *CustomerService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(customerModelIndex())
	return &CustomerService{collection}
}

func (p *CustomerService) FindAll() ([]*customer.Customer, error) {
	customers := []customerModel{}
	err := p.collection.Find(nil).All(&customers)
	allCustomer := []*customer.Customer{}
	for _, cst := range customers {
		allCustomer = append(allCustomer, cst.toRootCustomer())
	}
	return allCustomer, err

}

func (p *CustomerService) FindByAccountNumber(accountNumber string) (*customer.Customer, error) {
	customer := customerModel{}
	err := p.collection.Find(bson.M{"account_number": accountNumber}).One(&customer)
	return customer.toRootCustomer(), err

}

func (p *CustomerService) Update(accountNumber string, cst customer.Customer) (*mgo.ChangeInfo, error) {
	mci, err := p.collection.Upsert(bson.M{"account_number": accountNumber}, &cst)
	if err != nil {
		log.Print("Failed")
	}

	ci := &mgo.ChangeInfo{}
	if mci != nil {
		ci.Updated = mci.Updated
		ci.Removed = mci.Removed
		ci.UpsertedId = mci.UpsertedId
	}
	return ci, err
}
