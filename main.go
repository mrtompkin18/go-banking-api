package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"bank-api/handler"
	"bank-api/logger"
	"bank-api/repository"
	"bank-api/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initTimeZone()

	database := initDatabase()
	customerRepository := repository.NewCustomerRepository(database)
	customerService := service.NewCustomerCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	accountRepository := repository.NewAccountRepository(database)
	accountService := service.NewAccountService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)

	router := mux.NewRouter()

	/* Customer */
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	/* Account */
	router.HandleFunc("/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customerId:[0-9]+}/accounts", accountHandler.GetAccount).Methods(http.MethodGet)

	port := fmt.Sprintf(":%v", viper.GetString("server.port"))
	logger.Log.Info("Starting on port " + port + "...")

	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}

func initDatabase() *sqlx.DB {
	source := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.db-name"),
	)

	db, err := sqlx.Open("mysql", source)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)
	return db
}

func initConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}
