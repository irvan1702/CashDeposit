package mongo_test

import (
	"CashTest/customer"
	"CashTest/mongo"
	"log"
	"testing"
)

const (
	mongoUrl               = "localhost:27017"
	dbName                 = "customers_db"
	customerCollectionName = "customers"
)

func Test_CustomerService(t *testing.T) {
	t.Run("UpdateCustomer", UpdateCustomer)
	t.Run("FindByAccountNumber", FindByAccountNumber)
	t.Run("FindByAccountNumber", FindAllCustomer)
}

func UpdateCustomer(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable Connect to mongo : %s", err)
	}
	defer func() {
		session.DropDatabase(dbName)
		session.Close()
	}()
	customerService := mongo.NewCustomerService(session.Copy(), dbName, customerCollectionName)

	testCustomerName := "BABA"
	testAccountNumber := "940193"
	testCashDepost := 10000
	testCashHistory := []int{1000, 1000}
	testTotalCash := 10000

	customerTest := customer.Customer{
		CustomerName:       testCustomerName,
		AccountNumber:      testAccountNumber,
		CashDeposit:        testCashDepost,
		CashDepositHistory: testCashHistory,
		TotalCashDeposit:   testTotalCash,
	}

	info, err := customerService.Update(testAccountNumber, customerTest)
	if err != nil {
		t.Error("Unable to update/create user :%s", err)
	}

	if info.UpsertedId == nil || info.Updated != 0 {
		t.Error("Failed to update customer")
	}

	var result []customer.Customer
	session.GetCollection(dbName, customerCollectionName).Find(nil).All(&result)

	count := len(result)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `0`, got: `%i`", count)
	}

	if result[0].CustomerName != customerTest.CustomerName {
		t.Error("Incorrect Username. Expected `%s`, Got: `%s`", testCustomerName, result[0].CustomerName)
	}

}

func FindAllCustomer(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable Connect to mongo : %s", err)
	}
	defer func() {
		session.DropDatabase(dbName)
		session.Close()
	}()
	customerService := mongo.NewCustomerService(session.Copy(), dbName, customerCollectionName)

	testCustomerName := "Irvan"
	testAccountNumber := "940193"
	testCashDepost := 10000
	testCashHistory := []int{1000, 1000}
	testTotalCash := 10000

	customerTest := customer.Customer{
		CustomerName:       testCustomerName,
		AccountNumber:      testAccountNumber,
		CashDeposit:        testCashDepost,
		CashDepositHistory: testCashHistory,
		TotalCashDeposit:   testTotalCash,
	}

	info, err := customerService.Update(testAccountNumber, customerTest)
	if err != nil {
		t.Error("Unable to update/create user :%s", err)
	}

	if info.UpsertedId == nil || info.Updated != 0 {
		t.Error("Failed to update customer")
	}

	var result []customer.Customer
	session.GetCollection(dbName, customerCollectionName).Find(nil).All(&result)

	count := len(result)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got: `%i`", count)
	}

	if result[0].CustomerName != customerTest.CustomerName {
		t.Error("Incorrect Username. Expected `%s`, Got: `%s`", testCustomerName, result[0].CustomerName)
	}

}

func FindByAccountNumber(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable Connect to mongo : %s", err)
	}
	defer func() {
		session.DropDatabase(dbName)
		session.Close()
	}()
	customerService := mongo.NewCustomerService(session.Copy(), dbName, customerCollectionName)

	testCustomerName := "Irvan"
	testAccountNumber := "940193"
	testCashDepost := 10000
	testCashHistory := []int{1000, 1000}
	testTotalCash := 10000

	customerTest := customer.Customer{
		CustomerName:       testCustomerName,
		AccountNumber:      testAccountNumber,
		CashDeposit:        testCashDepost,
		CashDepositHistory: testCashHistory,
		TotalCashDeposit:   testTotalCash,
	}

	info, err := customerService.Update(testAccountNumber, customerTest)
	if err != nil {
		t.Error("Unable to update/create user :%s", err)
	}

	if info.UpsertedId == nil || info.Updated != 0 {
		t.Error("Failed to update customer")
	}

	result, err := customerService.FindByAccountNumber("940193")
	if err != nil {
		t.Error("Unable to find customer:%s", err)
	}

	actual := customer.Customer{}
	session.GetCollection(dbName, customerCollectionName).Find(nil).One(&actual)

	if result.CustomerName != actual.CustomerName {
		t.Error("Incorrect Username. Expected `%s`, Got: `%s`", result.CustomerName, actual.CustomerName)
	}

}
