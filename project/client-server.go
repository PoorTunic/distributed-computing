package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"project/dao"
	"project/web"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

var dataController *web.DataController

var server string
var port int

func main() {
	flag.StringVar(&server, "server", "localhost:", "SERVER")
	flag.IntVar(&port, "port", 8020, "PORT")
	flag.Parse()
	dataDao, err := dao.GetDao()
	if err != nil {
		fmt.Println(err)
	}
	dataController = web.NewDataController(dataDao)
	router := web.NewRouter(dataController)
	dataController.ServerIP = "http://localhost:" + strconv.Itoa(port)
	n := negroni.New()
	n.UseHandler(router)
	newPort := ":" + strconv.Itoa(port)
	fmt.Println("Running . . ." + dataController.ServerIP + " server")
	sendNotification()
	n.Run(newPort)
}

func sendNotification() {
	if port != 8020 {
		url := "http://localhost:8020/servers"
		log.Info("SENDING SERVER INFO TO " + url)
		serverValues := map[string]string{
			"ip": "http://localhost:" + strconv.Itoa(port),
		}
		serverJSON, _ := json.Marshal(serverValues)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(serverJSON))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Error(err.Error())
		}
		defer resp.Body.Close()
	}
}
