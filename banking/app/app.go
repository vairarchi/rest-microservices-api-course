package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vairarchi/rest-microservices-api-course/banking/confighelper"
	"github.com/vairarchi/rest-microservices-api-course/banking/domain"
	"github.com/vairarchi/rest-microservices-api-course/banking/service"
)

func Start() {

	confighelper.InitViper()

	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	fmt.Println("Server started on :", confighelper.GetConfig("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", confighelper.GetConfig("SERVER_ADDRESS"), confighelper.GetConfig("SERVER_PORT")), router))
}
