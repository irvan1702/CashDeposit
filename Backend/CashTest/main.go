package main

import (
	"CashTest/mongo"
	"CashTest/server"
	"html/template"
	"log"
)

var render *template.Template

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Fatalln("Unable to connect to mongodb")
	}

	defer ms.Close()
	cst := mongo.NewCustomerService(ms.Copy(), "customers_db", "customers")
	s := server.NewServer(cst)

	s.Start()
}
