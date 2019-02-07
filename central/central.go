package main

import (
	"central/dao"
	"central/web"
	"fmt"

	"github.com/urfave/negroni"
)

var dataController *web.DataController

var server string
var port int

func main() {
	dataDao, err := dao.GetDao()
	if err != nil {
		fmt.Println(err)
	}
	dataController = web.NewDataController(dataDao)
	router := web.NewRouter(dataController)
	n := negroni.New()
	n.UseHandler(router)
	fmt.Println("Running server")
	n.Run()
}
