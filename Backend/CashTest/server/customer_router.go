package server

import (
	"CashTest/customer"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type customerRouter struct {
	customerService customer.CustomerService
}

func NewCustomerRouter(cst customer.CustomerService, router *mux.Router) *mux.Router {
	customerRouter := customerRouter{cst}

	router.HandleFunc("/", customerRouter.findAllCustomer).Methods("GET")
	router.HandleFunc("/update", customerRouter.updateCustomer).Methods("POST")
	router.HandleFunc("/{accountNumber}", customerRouter.findCustomerByAccountNumber).Methods("GET")
	return router
}

func (cst *customerRouter) findAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, err := cst.customerService.FindAll()
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, customers)
}

func (cst *customerRouter) findCustomerByAccountNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountNumber := vars["accountNumber"]
	fmt.Println(accountNumber)
	customer, err := cst.customerService.FindByAccountNumber(accountNumber)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, customer)
}

func (cst *customerRouter) updateCustomer(w http.ResponseWriter, r *http.Request) {
	req, err := decodeCustomer(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	params := req.AccountNumber
	existCustomer, er := cst.customerService.FindByAccountNumber(params)
	if er != nil {
		req.CashDepositHistory = append(req.CashDepositHistory,req.CashDeposit)
		req.TotalCashDeposit = req.CashDeposit
		info, err := cst.customerService.Update(params, req)
		if err != nil {
			Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		Json(w, http.StatusOK, info)
	} else {
		existCustomer.AccountNumber = req.AccountNumber
		existCustomer.CashDeposit = req.CashDeposit
		existCustomer.CustomerName = req.CustomerName
		existCustomer.CashDepositHistory = append(existCustomer.CashDepositHistory, req.CashDeposit)
		existCustomer.TotalCashDeposit += existCustomer.CashDepositHistory[len(existCustomer.CashDepositHistory)-1]
		info, err := cst.customerService.Update(params, *existCustomer)
		if err != nil {
			Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		Json(w, http.StatusOK, info)
	}
}

func decodeCustomer(r *http.Request) (customer.Customer, error) {
	var u customer.Customer
	if r.Body == nil {
		return u, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	return u, err
}
