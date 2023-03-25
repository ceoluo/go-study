package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var EmployeesDB map[string]*Employee

func init() {
	EmployeesDB = map[string]*Employee{}
	EmployeesDB["Mike"] = &Employee{"1", "Mike", 35}
	EmployeesDB["Rose"] = &Employee{"2", "Rose", 30}
	EmployeesDB["zhongxiao"] = &Employee{"2", "zhongxiao", 23}
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/employees/:name", GetEmployeeByName)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetEmployeeByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeName := params.ByName("name")
	infoJson, err := queryEmployee(employeeName)
	if err != nil {
		_, _ = fmt.Fprintf(writer, "Not Found, error: %s", err)
	} else {
		_, _ = fmt.Fprint(writer, string(infoJson))
	}
}

func queryEmployee(name string) ([]byte, error) {
	if info, ok := EmployeesDB[name]; !ok{
		return nil, errors.New("the employee not in DB!")
	}else {
		if infoJson, err := json.Marshal(info); err!=nil{
			return nil, err
		}else {
			return infoJson, nil
		}
	}
}

func index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, _ = fmt.Fprint(writer, "Welcome!\n")
}
