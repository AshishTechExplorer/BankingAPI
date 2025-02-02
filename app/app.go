package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

func Start() {
	router := mux.NewRouter()
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepoDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ac := AccountHandlers{service.NewAccountService(accountRepoDb)}
	//	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))} //router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/account", ac.saveAccountDetails).Methods(http.MethodPost)
	//router.HandleFunc("/customers/{customer_id}", ch.getCustomerById).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

func getDbClient() *sqlx.DB {
	//dbUser := os.Getenv("DB_USER")
	dbUser := "root"
	dbPasswd := "root"
	dbAddr := "localhost"
	dbPort := "3306"
	dbName := "banking"
	//	dbPasswd := os.Getenv("DB_PASSWD")
	//dbAddr := os.Getenv("DB_ADDR")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
