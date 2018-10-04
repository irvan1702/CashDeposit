package mongo

import (
	"CashTest/customer"

	mgo "gopkg.in/mgo.v2"
)

type customerModel struct {
	CustomerName       string `bson:"customer_name"`
	AccountNumber      string `bson:"account_number"`
	CashDeposit        int    `bson:"cash_deposit"`
	CashDepositHistory []int  `bson:"cash_deposit_history"`
	TotalCashDeposit   int    `bson:"total_cash_deposit"`
}

func customerModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"account_number"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newCustomerModel(u *customer.Customer) *customerModel {
	return &customerModel{
		CustomerName:       u.CustomerName,
		AccountNumber:      u.AccountNumber,
		CashDeposit:        u.CashDeposit,
		CashDepositHistory: u.CashDepositHistory,
		TotalCashDeposit:   u.TotalCashDeposit,
	}
}

func (u *customerModel) toRootCustomer() *customer.Customer {
	return &customer.Customer{
		CustomerName:       u.CustomerName,
		AccountNumber:      u.AccountNumber,
		CashDeposit:        u.CashDeposit,
		CashDepositHistory: u.CashDepositHistory,
		TotalCashDeposit:   u.TotalCashDeposit,
	}
}
