package controller

import (
	"net/http"
	"summary/domain"
	"summary/dao"
	"strconv"
	"summary/infrastructure"
	"summary/utils"
	"io"
	"fmt"
)

var _logger utils.Logger = utils.NewFileLogger()

type CustomerHandler struct {
	dbRepo *dao.CustomerDbRepo
}

func NewCustomerHandler() *CustomerHandler {
	handler := infrastructure.NewMysqlHandler()
	customerRepo := dao.NewCustomerRepo(handler)
	return &CustomerHandler{dbRepo: customerRepo}
}

func (handler *CustomerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	_logger.Info(method + "\n")

	switch method {
	case "GET":
		handler.getCustomer(writer, request)
	case "POST":
		handler.addCustomer(request)
	case "PUT":
		break
	case "DELETE":
		break
	}
}

func (handler *CustomerHandler) getCustomer(writer http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.FormValue("id"))
	customer:=handler.dbRepo.FindById(id)
	io.WriteString(writer, fmt.Sprintf("customer id: %d\n", customer.Id))
	io.WriteString(writer, fmt.Sprintf("customer name: %v\n", customer.Name))
}

func (handler *CustomerHandler) addCustomer(req *http.Request) {
	id, _ := strconv.Atoi(req.FormValue("id"))
	name := req.FormValue("name")

	customer := domain.Customer{Id: id, Name: name}
	handler.dbRepo.AddCustomer(customer)
}
